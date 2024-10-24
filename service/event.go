package service

import (
	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/repository"
)

type EventService interface {
	GetEventsForChild(childID int, adultID int) ([]model.Event, error)
}

type SimpleEventService struct {
	Repository repository.Repository
}

func NewSimpleEventService(r repository.Repository) SimpleEventService {
	return SimpleEventService{Repository: r}
}

func (ses SimpleEventService) GetEventsForChild(childID int, adultID int) ([]model.Event, error) {
	var events []model.Event
	events, err := ses.Repository.GetEventsForChild(childID, adultID)
	if err != nil {
		return events, err
	}
	return events, err
}
