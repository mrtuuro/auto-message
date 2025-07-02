package autosend

import (
	"context"
	"time"
)

type Flusher interface {
	Flush(ctx context.Context, max int) error
}

type Scheduler struct {
	f      Flusher
	ticker *time.Ticker
	cancel context.CancelFunc
}

func NewScheduler(f Flusher) *Scheduler {
	return &Scheduler{
		f: f,
	}
}

func (s *Scheduler) Start() {
	if s.ticker != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.ticker = time.NewTicker(2 * time.Minute)

	go func() {
		for {
			select {
			case <-s.ticker.C:
				_ = s.f.Flush(ctx, 2)
			case <-ctx.Done():
				s.ticker.Stop()
				return
			}
		}
	}()
}

func (s *Scheduler) Stop() {
	if s.cancel != nil {
		s.cancel()
		s.ticker = nil
	}
}
