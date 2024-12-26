package handler

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

const sessionTimeout = 15 * time.Minute

// func hashPassword(password string) string {
// 	hash := sha256.Sum256([]byte(password))
// 	return base64.StdEncoding.EncodeToString(hash[:])
// }

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// func validateSession(db *sql.DB, token string) (string, error) {
// 	var username string
// 	var createdAt time.Time
// 	err := db.QueryRow(`SELECT username, created_at FROM session WHERE token = ?`, token).Scan(&username, &createdAt)
// 	if err != nil {
// 		return "", err
// 	}
// 	if time.Since(createdAt) > sessionTimeout {
// 		_, err := db.Exec(`DELETE FROM sessions WHERE token = ?`, token)
// 		if err != nil {
// 			log.Printf("Failed to delete expired session: %v", err)
// 		}
// 		return "", fmt.Errorf("session expired")
// 	}
// 	return username, nil
// }

func ValidateSession(db *sql.DB, r *http.Request) error {
	var userID string
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return err
	}
	token := cookie.Value
	err = db.QueryRow("SELECT username FROM session WHERE token = ?", token).Scan(&userID)
	if err != nil {
		return err
	}
	return nil
}

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
		token, err := generateToken()
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
