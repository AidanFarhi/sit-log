package controller

import (
	"net/http"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

type EventController struct {
	Service service.EventService
}

func NewEventController(s service.EventService) EventController {
	return EventController{Service: s}
}

func (ec EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	newEvent := model.NewEvent{}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusInternalServerError)
		return
	}
	// childId, err := strconv.Atoi(r.Form.Get("childId"))
	// if err != nil {
	// 	http.Error(w, "invalid childId", http.StatusInternalServerError)
	// 	return
	// }
	// newEvent.ChildID = childId
	newEvent.Type = r.Form.Get("eventType")
	newEvent.Description = r.Form.Get("description")
	newEvent.StartTime = r.Form.Get("startTime")
	newEvent.EndTime = r.Form.Get("endTime")
	switch {
	case newEvent.Type == "":
		http.Error(w, "event type cannot be empty", http.StatusInternalServerError)
		return
	case newEvent.StartTime == "":
		http.Error(w, "start time cannot be empty", http.StatusInternalServerError)
		return
	case newEvent.EndTime == "":
		http.Error(w, "end time cannot be empty", http.StatusInternalServerError)
		return
	}
	err = ec.Service.CreateEvent(newEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
