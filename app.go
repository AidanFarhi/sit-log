package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AidanFarhi/sitlog/handler"
	"github.com/AidanFarhi/sitlog/model"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "sitlog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	templates := model.NewTemplates()

	fs := http.FileServer(http.Dir("web"))
	mux := http.NewServeMux()

	mux.Handle("/web/", http.StripPrefix("/web/", fs))

	mux.HandleFunc("/", handler.IndexHandler(db, templates))
	mux.HandleFunc("/login", handler.LoginHandler(db, templates))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
