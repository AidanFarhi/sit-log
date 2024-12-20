package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AidanFarhi/sitlog/controller"
	"github.com/AidanFarhi/sitlog/handler"
	"github.com/AidanFarhi/sitlog/model"
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
	if err != nil {
		log.Fatal(err)
	}

	templates := model.NewTemplates()
	fs := http.FileServer(http.Dir("static"))
	pageData := model.NewPageData()

	eventRepo := repository.NewSimpleEventRepository(db)
	eventService := service.NewSimpleEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handler.Index(pageData, templates, eventService))
	mux.HandleFunc("GET /api/v1/events/{childId}/{adultId}", eventController.GetEventsForChild)
	mux.HandleFunc("POST /api/v1/events/create", eventController.CreateEvent)
	mux.HandleFunc("POST /login", handler.LoginHandler(db))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
