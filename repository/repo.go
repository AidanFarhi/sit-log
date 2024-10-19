package repository

import "github.com/AidanFarhi/sitlog/model"

type EventRepository interface {
	GetEvent(ID int) (model.Event, error)
	GetEventsForChildID(childID int) ([]model.Event, error)
	CreateEvent(event model.Event) error
}
