package runner

import (
	"time"
)

type Runner struct {
	interval time.Duration
	quitter  chan struct{}
}

func New(interval time.Duration) *Runner {
	return &Runner{
		interval: interval,
		quitter:  make(chan struct{}),
	}
}

func (r *Runner) Run(tick func()) {
	ticker := time.NewTicker(r.interval)
	for {
		select {
		case <-ticker.C:
			tick()

		case <-r.quitter:
			return
		}
	}
}

func (r *Runner) Close() {
	r.quitter <- struct{}{}
}
