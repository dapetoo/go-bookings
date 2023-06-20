package main

import (
	"github.com/dapetoo/go-bookings/internal/config"
	"github.com/dapetoo/go-bookings/internal/handlers"
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
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	//mux.Use(middleware.CleanPath)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//HTTP Server
	mux.Get("/", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.Home)))
	mux.Get("/about", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.About)))
	mux.Get("/contact", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.Contact)))
	mux.Get("/service", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.Service)))
	mux.Get("/generals-quarters", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.Generals)))
	mux.Get("/majors-suite", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.Majors)))

	mux.Get("/search-availability", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.Availability)))
	mux.Post("/search-availability", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.PostAvailability)))
	mux.Post("/search-availability-json", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.AvailabilityJSON)))
	mux.Get("/choose-room/{id}", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.ChooseRoom)))
	mux.Get("/book-room", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.BookRoom)))

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/user/login", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.ShowLogin)))
	mux.Post("/user/login", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.PostShowLogin)))
	mux.Get("/user/logout", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.Logout)))

	// File server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.AdminDashboard)))
		mux.Get("/reservations-new", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.AdminNewReservations)))
		mux.Get("/reservations-all", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.AdminAllReservations)))
		mux.Get("/reservations-calendar", sentryHandler.HandleFunc(EnhanceSentryEvent(handlers.Repo.AdminReservationsCalendar)))
	})

	return mux
}
