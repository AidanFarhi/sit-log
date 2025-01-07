package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

func CreateEventHandler(db *sql.DB, t model.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: first check if user session is valid
		err := service.CreateEvent(db, r)
		if err != nil {
			fmt.Println("error creating event")
			http.Error(w, "error creating event", http.StatusInternalServerError)
		}
		// TODO: pull the child ID from service and use session to get adult id
		events, err := service.GetEventsForChild(db, 2, 2)
		if err != nil {
			fmt.Println("error getting events")
			http.Error(w, "error getting events", http.StatusInternalServerError)
			return
		}
		t.Templates.ExecuteTemplate(w, "event_list", events)
	}
}
