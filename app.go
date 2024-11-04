package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/AidanFarhi/sitlog/controller"
	"github.com/AidanFarhi/sitlog/repository"
	"github.com/AidanFarhi/sitlog/service"
	_ "github.com/mattn/go-sqlite3"
)

type Templates struct {
	Templates *template.Template
}

func NewTemplates() Templates {
	return Templates{
		Templates: template.Must(template.ParseGlob("static/views/*.html")),
	}
}

func main() {

	db, err := sql.Open("sqlite3", "sitlog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	templates := NewTemplates()

	eventRepo := repository.NewSimpleEventRepository(db)
	eventService := service.NewSimpleEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		templates.Templates.ExecuteTemplate(w, "index", nil)
	})
	mux.HandleFunc("GET /api/v1/events/{childId}/{adultId}", eventController.GetEventsForChild)
	mux.HandleFunc("POST /api/v1/events/create", eventController.CreateEvent)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
