package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	fmt.Println("childID:", childID, "adultID:", adultID)
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
