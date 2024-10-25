package model

type Event struct {
	TimeStamp   string
	Type        string
	Description string
	StartTime   string
	EndTime     string
	Duration    string
}

type NewEvent struct {
	ChildID     int
	Type        string
	Description string
	StartTime   string
	EndTime     string
	Duration    string
}
