package main

import (
	"net/http"

	"github.com/catehulu/factorio-calculator/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Index)
	mux.Post("/calculate", handlers.Calculate)
	return mux
}
