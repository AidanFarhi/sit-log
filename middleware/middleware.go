package middleware

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/AidanFarhi/sitlog/model"
)

type MiddleWare struct {
	Templates model.Templates
	DB        *sql.DB
}

func (mw MiddleWare) SessionValidation(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		fmt.Println("cookie", cookie)
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				mw.Templates.Templates.ExecuteTemplate(w, "index", model.PageData{IsLoggedIn: false, Events: []model.Event{}})
			} else {
				w.WriteHeader(http.StatusBadRequest)
				mw.Templates.Templates.ExecuteTemplate(w, "index", model.PageData{IsLoggedIn: false, Events: []model.Event{}})
			}
			return
		}
		token := cookie.Value
		if !mw.isValidToken(token) {
			w.WriteHeader(http.StatusUnauthorized)
			mw.Templates.Templates.ExecuteTemplate(w, "index", model.PageData{IsLoggedIn: false, Events: []model.Event{}})
			return
		}
		next(w, r)
	})
}

func (mw MiddleWare) isValidToken(token string) bool {
	var userID string
	err := mw.DB.QueryRow("SELECT username FROM session WHERE token = ?", token).Scan(&userID)
	fmt.Println(userID)
	return err == nil
}
