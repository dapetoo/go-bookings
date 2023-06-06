package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/dapetoo/go-bookings/pkg/config"
	"github.com/dapetoo/go-bookings/pkg/handlers"
	"github.com/dapetoo/go-bookings/pkg/render"
	"github.com/getsentry/sentry-go"
	"github.com/rollbar/rollbar-go"
	"log"
	"net/http"
	"time"
)

type handler struct{}

func (h *handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			scope.SetExtra("unwantedQuery", "someQueryDataMaybe")
			hub.CaptureMessage("User provided unwanted query string, but we recovered just fine")
		})
	}
	rw.WriteHeader(http.StatusOK)
}

func EnhanceSentryEvent(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		handler(rw, r)
	}
}

func doSomething() {
	var timer *time.Timer = nil
	timer.Reset(10) // this will panic
}

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	rollbar.SetToken("491f9cbe3ace480dabe3544419b18221")
	rollbar.SetEnvironment("production")                 // defaults to "development"
	rollbar.SetCodeVersion("v2")                         // optional Git hash/branch/tag (required for GitHub integration)
	rollbar.SetServerHost("web.1")                       // optional override; defaults to hostname
	rollbar.SetServerRoot("github.com/heroku/myproject") // path of project (required for GitHub integration and non-project stacktrace collapsing)

	rollbar.Info("Message body goes here")
	rollbar.WrapAndWait(doSomething)
	rollbar.Log(rollbar.INFO, "Message body goes here")
	rollbar.Wait()

	// call rollbar.Close() before the application exits to flush error message queue
	rollbar.Close()

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://5e80ef8860d544c1ac1db0a1e4b55328@o4505127968374784.ingest.sentry.io/4505257065316352",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
		Debug:            true,
		Environment:      "development",
		// Enable this to see the full event
		// Timeout for the event delivery requests.
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	sentry.CaptureMessage("It works!")
	defer sentry.Flush(2 * time.Second)

	//Change this to true when in production
	app.InProduction = false

	//Session Management
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //Set to true in production

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", sentryHandler.HandleFunc(enhanceSentryEvent(handlers.Repo.Home)))
	//http.HandleFunc("/about", sentryHandler.HandleFunc(enhanceSentryEvent(handlers.Repo.About)))
	//http.HandleFunc("/service", sentryHandler.HandleFunc(enhanceSentryEvent(handlers.Repo.Service)))
	///http.HandleFunc("/contact", sentryHandler.HandleFunc(enhanceSentryEvent(handlers.Repo.Contact)))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

	//_ = http.ListenAndServe(portNumber, nil)

}
