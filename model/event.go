package model

import "time"

type Event struct {
	TimeStamp     time.Time
	Type          string
	Description   string
	StartTime     time.Time
	EndTime       time.Time
	EventDuration time.Duration
}
