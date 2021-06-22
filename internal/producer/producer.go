package producer

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

type EventType int

const (
	Created EventType = iota
	Updated
	Deleted
)

type Event struct {
	Type EventType
	Body map[string]interface{}
}

type Producer interface {
	Send(Event) error
	Close() error
}

type ProducerBuffered interface {
	Producer
	C() <-chan Event
	Done() <-chan struct{}
}

type producer struct {
	topic string
	p     sarama.SyncProducer
}

func NewProducer(addrs []string, topic string) (Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducer(addrs, cfg)

	if err != nil {
		return nil, err
	}

	return &producer{
		topic: topic,
		p:     p,
	}, nil
}

type nullProducer struct {
}

func NewNullProducer() Producer {
	return &nullProducer{}
}

type buffered struct {
	done chan struct{}
	ch   chan Event
	p    Producer
}

func (p *producer) Send(e Event) error {
	bs, err := json.Marshal(e)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(bs),
	}

	_, _, err = p.p.SendMessage(msg)

	return err
}

func (p *producer) Close() error {
	return p.p.Close()
}

func (b *buffered) Send(e Event) error {
	select {
	case <-b.done:
		close(b.ch)
		break
	case b.ch <- e:
	}

	return nil
}

func (b *buffered) Close() error {
	b.done <- struct{}{}
	return b.p.Close()
}

func (b *buffered) C() <-chan Event {
	return b.ch
}

func (b *buffered) Done() <-chan struct{} {
	return b.done
}

func NewBuffered(p Producer, buffSize int) (ProducerBuffered, error) {
	return &buffered{
		done: make(chan struct{}),
		ch:   make(chan Event, buffSize),
		p:    p,
	}, nil
}

func (*nullProducer) Send(Event) error {
	return nil
}

func (*nullProducer) Close() error {
	return nil
}
