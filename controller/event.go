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

func (ec EventController) GetEventsForAdult(w http.ResponseWriter, r *http.Request) {
	fmt.Println("query params", r.URL.Query())
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Println(id)
	if err != nil {
		http.Error(w, "Invalid Adult ID", http.StatusBadRequest)
		return
	}
	events, err := ec.Service.GetEventsForAdult(id)
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
