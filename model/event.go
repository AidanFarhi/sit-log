package model

import "time"

type Event struct {
	AdultID       int
	TimeStamp     string
	Type          string
	Description   string
	StartTime     string
	EndTime       string
	EventDuration time.Duration
}
