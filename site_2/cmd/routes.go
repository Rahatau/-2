package main

import (
	"app/pkg/handlers"
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.HomePage))

	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Get("/", handlers.HomePage)
	return mux
}
