package model

import "time"

type Event struct {
	ID            int
	AdultID       int
	ChildID       int
	TimeStamp     string
	Type          string
	Description   string
	StartTime     string
	EndTime       string
	EventDuration time.Duration
}
