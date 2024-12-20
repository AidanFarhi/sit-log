package model

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type PageData struct {
	IsLoggedIn bool
	Events     []Event
}

func NewPageData() PageData {
	return PageData{false, []Event{}}
}

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
