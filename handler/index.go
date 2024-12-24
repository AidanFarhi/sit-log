package handler

import (
	"log"
	"net/http"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

func IndexHandler(t model.Templates, es service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := model.PageData{}
		if r.Method == http.MethodGet {
			events, _ := es.GetEventsForChild(2, 2)
			pageData.Events = events
			err := t.Templates.ExecuteTemplate(w, "index", pageData)
			if err != nil {
				log.Printf("Error executing template: %v", err)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
