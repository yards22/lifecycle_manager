package runner

import (
	"time"
)

type Runner struct {
	interval time.Duration
	quitter  chan struct{}
}

const TUsersFrequency = 12
const TPostsFrequency = 6
const RUsersFrequency = 12
const RPostsFrequency = 6
const TCleanerFrequency = 24
const RatingFrequency = 7 * 24

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
