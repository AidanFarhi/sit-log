package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AidanFarhi/sitlog/controller"
	"github.com/AidanFarhi/sitlog/repository"
	"github.com/AidanFarhi/sitlog/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "sitlog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewSimpleEventRepository(db)
	service := service.NewSimpleEventService(repo)
	controller := controller.NewEventController(service)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/events/adult/{id}", controller.GetEventsForAdult)
	mux.HandleFunc("GET /api/v1/events/child/{id}", controller.GetEventsForChild)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
