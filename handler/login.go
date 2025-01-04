package handler

import (
	"database/sql"
	"net/http"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

func LoginHandler(db *sql.DB, t model.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		err := service.ValidateLogin(db, r)
		if err != nil {
			pageData := model.PageData{IsLoggedIn: false, Error: true, Events: []model.Event{}}
			t.Templates.ExecuteTemplate(w, "index", pageData)
			return
		}
		token, err := service.CreateNewSession(db, r)
		if err != nil {
			http.Error(w, "error creating session", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: token,
			Path:  "/",
		})
		pageData := model.PageData{IsLoggedIn: true, Error: false, Events: []model.Event{}}
		events, err := service.GetEventsForChild(db, 2, 2)
		if err != nil {
			http.Error(w, "error getting events", http.StatusInternalServerError)
			return
		}
		pageData.Events = events
		t.Templates.ExecuteTemplate(w, "index", pageData)
	}
}
