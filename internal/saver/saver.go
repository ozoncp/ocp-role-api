package saver

import (
	"context"
	"sync"
	"time"

	"github.com/ozoncp/ocp-role-api/internal/flusher"
	"github.com/ozoncp/ocp-role-api/internal/model"
	"github.com/ozoncp/ocp-role-api/internal/ticker"
)

type Saver interface {
	Save(entity *model.Role)
	Close()
}

type Option func(*saver)

func New(capacity uint, flusher flusher.Flusher, opt Option) Saver {
	s := &saver{
		flusher: flusher,
		data:    make([]*model.Role, 0, capacity),
		doneCh:  make(chan struct{}),
	}

	opt(s)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-s.doneCh
		cancel()
	}()

	go func() {
		defer cancel()
		tickerCh := s.ticker.C()

		for {
			select {
			case <-tickerCh:
				var (
					data []*model.Role
					end  int
				)

				func() {
					s.mu.Lock()
					defer s.mu.Unlock()

					if len(s.data) == 0 {
						return
					}

					end = s.end
					data = s.data

					s.data = make([]*model.Role, 0, cap(s.data))
					s.end = 0
				}()

				if end == 0 {
					continue
				}

				data = s.flusher.Flush(ctx, getOrdered(data, end))
				if data != nil {
					s.mu.Lock()
					s.data = data
					s.end = len(data)
					s.mu.Unlock()
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return s
}

func WithTicker(ticker ticker.Ticker) Option {
	return func(s *saver) {
		s.ticker = ticker
	}
}

func WithPeriod(period time.Duration) Option {
	return func(s *saver) {
		s.ticker = ticker.New(period)
	}
}

type saver struct {
	end     int
	flusher flusher.Flusher
	data    []*model.Role
	ticker  ticker.Ticker
	doneCh  chan struct{}
	mu      sync.Mutex
}

func (s *saver) Save(role *model.Role) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.data) < cap(s.data) {
		s.data = append(s.data, role)
		s.end++
	} else {
		s.data[s.end%len(s.data)] = role
		s.end = (s.end + 1) % len(s.data)
	}
}

func (s *saver) Close() {
	s.ticker.Stop()
	s.doneCh <- struct{}{}

	if len(s.data) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		s.data = s.flusher.Flush(ctx, getOrdered(s.data, s.end))
		s.end = len(s.data)
	}
}

func getOrdered(data []*model.Role, end int) []*model.Role {
	start := 0
	if len(data) == cap(data) {
		start = end % len(data)
	}

	if start >= end {
		ordered := make([]*model.Role, 0, cap(data))
		ordered = append(ordered, data[start:]...)
		ordered = append(ordered, data[:end]...)
		return ordered
	} else {
		return data
	}
}
