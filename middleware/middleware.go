package middleware

import (
	"database/sql"
	"net/http"
)

type MiddleWare struct {
	db *sql.DB
}

func (mw MiddleWare) SessionValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Missing session token", http.StatusUnauthorized)
			} else {
				http.Error(w, "Error retrieving session token", http.StatusBadRequest)
			}
			return
		}
		token := cookie.Value
		if !mw.isValidToken(token) {
			http.Error(w, "Invalid or expired session token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (mw MiddleWare) isValidToken(token string) bool {
	var userID string
	err := mw.db.QueryRow("SELECT user_id FROM sessions WHERE token = ?", token).Scan(&userID)
	return err == nil
}
