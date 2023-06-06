package main

import (
	"github.com/dapetoo/go-bookings/pkg/config"
	"github.com/dapetoo/go-bookings/pkg/handlers"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	// Create an instance of sentry
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//HTTP Server
	mux.Get("/", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.Home)))
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/service", handlers.Repo.Service)
	mux.Get("/make-reservation", handlers.Repo.About)
	mux.Get("/majors-suite", handlers.Repo.Contact)
	mux.Get("/reservation", handlers.Repo.Service)
	mux.Get("/generals-quarters", handlers.Repo.Service)

	// File server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
