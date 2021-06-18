package ticker

import "time"

type Ticker interface {
	Stop()
	Reset(d time.Duration)
	C() <-chan time.Time
}

type ticker struct {
	t *time.Ticker
}

func (t *ticker) Stop() {
	t.t.Stop()
}

func (t *ticker) Reset(d time.Duration) {
	t.t.Reset(d)
}

func (t *ticker) C() <-chan time.Time {
	return t.t.C
}

func New(d time.Duration) Ticker {
	return &ticker{time.NewTicker(d)}
}
