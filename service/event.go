package service

import (
	"fmt"
	"time"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/repository"
)

type EventService interface {
	GetEventsForChild(childID int, adultID int) ([]model.Event, error)
	CreateEvent(newEvent model.NewEvent) error
}

type SimpleEventService struct {
	Repository repository.EventRepository
}

func NewSimpleEventService(r repository.EventRepository) SimpleEventService {
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

func (ses SimpleEventService) CreateEvent(newEvent model.NewEvent) error {
	duration, err := calculateDuration(newEvent.StartTime, newEvent.EndTime)
	if err != nil {
		return err
	}
	newEvent.Duration = duration
	err = ses.Repository.CreateEvent(newEvent)
	if err != nil {
		return err
	}
	return nil
}

func calculateDuration(startTime string, endTime string) (string, error) {
	layout := "15:04:05"
	start, err := time.Parse(layout, startTime)
	if err != nil {
		return "", fmt.Errorf("invalid start time format: %v", err)
	}
	end, err := time.Parse(layout, endTime)
	if err != nil {
		return "", fmt.Errorf("invalid end time format: %v", err)
	}
	duration := end.Sub(start)
	if duration < 0 {
		duration += 24 * time.Hour
	}
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
}
