package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LusBlack/gollum/pkg/config"
	"github.com/LusBlack/gollum/pkg/handlers"
	"github.com/LusBlack/gollum/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
