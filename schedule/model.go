package schedule

import "time"

type schedule struct {
	to          []string
	template    string
	scheduledTo time.Time
}
