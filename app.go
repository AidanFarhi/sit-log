package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/AidanFarhi/sitlog/controller"
	"github.com/AidanFarhi/sitlog/model"
	"github.com/AidanFarhi/sitlog/repository"
	"github.com/AidanFarhi/sitlog/service"
	_ "github.com/mattn/go-sqlite3"
)

type Templates struct {
	Templates *template.Template
}

func NewTemplates() Templates {
	tmpl := template.New("")
	err := filepath.Walk("static/views", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			_, err := tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	return Templates{
		Templates: tmpl,
	}
}

type PageData struct {
	IsLoggedIn bool
	Events     []model.Event
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
	fs := http.FileServer(http.Dir("static"))
	pageData := PageData{false, []model.Event{}}

	eventRepo := repository.NewSimpleEventRepository(db)
	eventService := service.NewSimpleEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// todo: how do we make these values come from the client?
			events, _ := eventService.GetEventsForChild(2, 2)
			pageData.Events = events
			err := templates.Templates.ExecuteTemplate(w, "index", pageData)
			if err != nil {
				log.Printf("Error executing template: %v", err)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("GET /api/v1/events/{childId}/{adultId}", eventController.GetEventsForChild)
	mux.HandleFunc("POST /api/v1/events/create", eventController.CreateEvent)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
