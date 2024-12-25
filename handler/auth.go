package handler

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
)

const sessionTimeout = 15 * time.Minute

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func generateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func validateSession(db *sql.DB, token string) (string, error) {
	var username string
	var createdAt time.Time
	err := db.QueryRow(`SELECT username, created_at FROM sessions WHERE token = ?`, token).Scan(&username, &createdAt)
	if err != nil {
		return "", err
	}
	if time.Since(createdAt) > sessionTimeout {
		_, err := db.Exec(`DELETE FROM sessions WHERE token = ?`, token)
		if err != nil {
			log.Printf("Failed to delete expired session: %v", err)
		}
		return "", fmt.Errorf("session expired")
	}
	return username, nil
}

func IsSessionTokenValid(token string, db *sql.DB) bool {
	var userID string
	err := db.QueryRow("SELECT username FROM session WHERE token = ?", token).Scan(&userID)
	return err == nil
}

func LoginHandler(db *sql.DB, t model.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		var storedPassword string
		err := db.QueryRow(`SELECT password FROM user WHERE username = ?`, username).Scan(&storedPassword)
		// if err != nil || storedPassword != hashPassword(password) {
		// 	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		// 	return
		// }
		if err != nil || storedPassword != password {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		// token := generateToken()
		// _, err = db.Exec(`INSERT INTO session (token, username, created_at) VALUES (?, ?, ?)`, token, username, time.Now())
		// if err != nil {
		// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
		// 	return
		// }
		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: "def456sessiontoken",
			Path:  "/",
		})
		pageData := model.PageData{IsLoggedIn: true, Events: []model.Event{}}
		events, _ := service.GetEventsForChild(db, 2, 2)
		pageData.Events = events
		err = t.Templates.ExecuteTemplate(w, "index", pageData)
		if err != nil {
			http.Error(w, "error executing template", http.StatusInternalServerError)
		}
	}
}
