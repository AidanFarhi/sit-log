package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AidanFarhi/sitlog/handler"
	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "sitlog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	templates := model.NewTemplates()
	service := service.EventService{DB: db}

	fs := http.FileServer(http.Dir("web"))
	mux := http.NewServeMux()

	mux.Handle("/web/", http.StripPrefix("/web/", fs))
	mux.HandleFunc("/", handler.IndexHandler(templates, service))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
