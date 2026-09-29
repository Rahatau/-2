package main

import (
	"fmt"
	"log"
	"net/http"

	"app/pkg/config"
	"app/pkg/handlers"
	"app/render"
)

func main() {
	var app config.AppConfig

	tmpl, err := render.CreateTemplatesCache()

	if err != nil {
		log.Println("Ошибка в созданий кэша")
	}

	app.TemplatesCache = tmpl

	render.NewTemplates(&app)

	handlers.NewRepo(&app)

	fmt.Println("Server opened!")

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	server.ListenAndServe()
}
