package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	// create a new router that routes HTTP requests to different handlers
	mux := chi.NewRouter()

	// make the router use a new CORS handler with specific options
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// make the router use middleware, i.e. Heartbeat with "/ping" as it's route
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/send", app.SendMail)

	return mux
}
