package handler

import (
	"log"
	"net/http"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

func IndexHandler(pageData model.PageData, templates model.Templates, eventService service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// todo: how do we make these values come from the client?
			events, _ := eventService.GetEventsForChild(2, 2)
			pageData.Events = events
			err := templates.Templates.ExecuteTemplate(w, "index", pageData)
			if err != nil {
				log.Printf("Error executing template: %v", err)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
