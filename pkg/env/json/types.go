package json

import "time"

type RunnerInterval struct {
	Seconds time.Duration `json:"sec"`
	Minutes time.Duration `json:"min"`
	Hours   time.Duration `json:"hrs"`
}

type ManagerConfigs struct {
	Interval RunnerInterval
}
