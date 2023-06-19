package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func EnhanceSentryEvent(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		handler(rw, r)
	}
}

func initSentry() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           os.Getenv("SENTRY_DSN"),
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
		Debug:            false,
		Environment:      "development",
		// Enable this to see the full event
		// Timeout for the event delivery requests.
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	sentry.CaptureMessage("It works!")
	defer sentry.Flush(2 * time.Second)
}
