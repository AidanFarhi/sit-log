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

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		var storedPassword string
		err := db.QueryRow(`SELECT password FROM users WHERE username = ?`, username).Scan(&storedPassword)
		// if err != nil || storedPassword != hashPassword(password) {
		// 	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		// 	return
		// }
		if err != nil || storedPassword != password {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		token := generateToken()
		_, err = db.Exec(`INSERT INTO sessions (token, username, created_at) VALUES (?, ?, ?)`, token, username, time.Now())
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: token,
			Path:  "/",
		})
		fmt.Fprintf(w, "Login successful!")
	}
}
