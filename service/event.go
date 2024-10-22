package service

import (
	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/repository"
)

type EventService interface {
	GetEvent(ID int) (model.Event, error)
	GetEventsForAdult(adultID int) ([]model.Event, error)
	GetEventsForChild(childID int) ([]model.Event, error)
	CreateEvent(event model.Event) error
}

type SimpleEventService struct {
	Repository repository.EventRepository
}

func NewSimpleEventService(r repository.EventRepository) SimpleEventService {
	return SimpleEventService{Repository: r}
}

func (ses SimpleEventService) GetEvent(ID int) (model.Event, error) {
	var event model.Event
	event, err := ses.Repository.GetEvent(ID)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (ses SimpleEventService) GetEventsForAdult(adultID int) ([]model.Event, error) {
	var events []model.Event
	events, err := ses.Repository.GetEventsForAdult(adultID)
	if err != nil {
		return events, err
	}
	return events, err
}

func (ses SimpleEventService) GetEventsForChild(childID int) ([]model.Event, error) {
	var events []model.Event
	events, err := ses.Repository.GetEventsForAdult(childID)
	if err != nil {
		return events, err
	}
	return events, err
}

func (ses SimpleEventService) CreateEvent(event model.Event) error {
	err := ses.Repository.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil
}
