package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

func LoginHandler(db *sql.DB, t model.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		var storedPassword string
		err := db.QueryRow(`SELECT password FROM user WHERE username = ?`, username).Scan(&storedPassword)
		if err != nil || storedPassword != password {
			pageData := model.PageData{IsLoggedIn: false, Error: true, Events: []model.Event{}}
			err = t.Templates.ExecuteTemplate(w, "index", pageData)
			if err != nil {
				http.Error(w, "error executing template", http.StatusInternalServerError)
				return
			}
			return
		}
		token, err := service.GenerateToken()
		if err != nil {
			http.Error(w, "error generating token", http.StatusInternalServerError)
			return
		}
		_, err = db.Exec(`INSERT INTO session (token, username, created_at) VALUES (?, ?, ?)`, token, username, time.Now())
		if err != nil {
			http.Error(w, "error generating token", http.StatusInternalServerError)
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
		err = t.Templates.ExecuteTemplate(w, "index", pageData)
		if err != nil {
			http.Error(w, "error executing template", http.StatusInternalServerError)
			return
		}
	}
}
