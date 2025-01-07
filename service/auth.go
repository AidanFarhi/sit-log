package service

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"time"
)

const sessionTimeout = 15 * time.Minute

// func hashPassword(password string) string {
// 	hash := sha256.Sum256([]byte(password))
// 	return base64.StdEncoding.EncodeToString(hash[:])
// }

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

func CreateNewSession(db *sql.DB, r *http.Request) (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", err
	}
	username := r.FormValue("username")
	_, err = db.Exec(`DELETE FROM session WHERE username = ?`, username)
	if err != nil {
		return "", err
	}
	_, err = db.Exec(`INSERT INTO session (token, username, created_at) VALUES (?, ?, ?)`, token, username, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func ValidateLogin(db *sql.DB, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")
	var storedPassword string
	err := db.QueryRow(`SELECT password FROM user WHERE username = ?`, username).Scan(&storedPassword)
	if err != nil || storedPassword != password {
		return errors.New("error validating password")
	}
	return nil
}

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
