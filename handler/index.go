package handler

import (
	"database/sql"
	"net/http"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

func IndexHandler(db *sql.DB, t model.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pageData := model.PageData{IsLoggedIn: false, Error: false, Events: []model.Event{}}
			cookie, err := r.Cookie("session_token")
			if err != nil {
				pageData.IsLoggedIn = false
				t.Templates.ExecuteTemplate(w, "index", pageData)
				return
			}
			token := cookie.Value
			if !IsSessionTokenValid(token, db) {
				pageData.IsLoggedIn = false
				t.Templates.ExecuteTemplate(w, "index", pageData)
				return
			}
			events, err := service.GetEventsForChild(db, 2, 2)
			if err != nil {
				http.Error(w, "error calling event service", http.StatusInternalServerError)
				return
			}
			pageData.IsLoggedIn = true
			pageData.Events = events
			err = t.Templates.ExecuteTemplate(w, "index", pageData)
			if err != nil {
				http.Error(w, "error executing template", http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
