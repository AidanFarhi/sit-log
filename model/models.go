package model

type Adult struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Child struct {
	ID       int
	ParentID int
	Name     string
}

type Event struct {
	ID            int
	AdultID       int
	ChildID       int
	TimeStamp     string
	Type          string
	Description   string
	StartTime     string
	EndTime       string
	EventDuration string
}
