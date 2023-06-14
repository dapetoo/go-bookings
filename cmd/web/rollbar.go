package main

import (
	"github.com/rollbar/rollbar-go"
	"time"
)

func setupRollbar() {
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
}

func doSomething() {
	var timer *time.Timer = nil
	timer.Reset(10) // this will panic
}
