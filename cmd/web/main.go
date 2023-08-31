package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kpauletti/gwa/pkg/config"
	"github.com/kpauletti/gwa/pkg/handlers"
	"github.com/kpauletti/gwa/pkg/render"
)

const port = ":3000"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting server on port %v \n", port)
	http.ListenAndServe(port, nil)
}
