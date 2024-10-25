package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

type EventController struct {
	Service service.EventService
}

func NewEventController(s service.EventService) EventController {
	return EventController{Service: s}
}

func (ec EventController) GetEventsForChild(w http.ResponseWriter, r *http.Request) {
	childID, err := strconv.Atoi(r.PathValue("childId"))
	if err != nil {
		http.Error(w, "Invalid Child ID", http.StatusBadRequest)
		return
	}
	adultID, err := strconv.Atoi(r.PathValue("adultId"))
	if err != nil {
		http.Error(w, "Invalid Adult ID", http.StatusBadRequest)
		return
	}
	events, err := ec.Service.GetEventsForChild(childID, adultID)
	if err != nil {
		http.Error(w, "Error Getting Events", http.StatusInternalServerError)
		return
	}
	listItems := strings.Builder{}
	listItems.WriteString("<ul>")
	for _, e := range events {
		listItems.WriteString("<li>" + e.Type + " | " + e.Description + "</li>")
	}
	listItems.WriteString("</ul>")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(listItems.String()))
}

func (ec EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	newEvent := model.NewEvent{}
	childId, err := strconv.Atoi(r.Form.Get("childId"))
	if err != nil {
		http.Error(w, "invalid childId", http.StatusInternalServerError)
	}
	newEvent.ChildID = childId
	newEvent.Type = r.Form.Get("eventType")
	newEvent.Description = r.Form.Get("description")
	newEvent.StartTime = r.Form.Get("startTime")
	newEvent.EndTime = r.Form.Get("endTime")
	switch {
	case newEvent.Type == "":
		http.Error(w, "event type cannot be empty", http.StatusInternalServerError)
	case newEvent.StartTime == "":
		http.Error(w, "start time cannot be empty", http.StatusInternalServerError)
	case newEvent.EndTime == "":
		http.Error(w, "end time cannot be empty", http.StatusInternalServerError)
	}
	err = ec.Service.CreateEvent(newEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("created event"))
}
