package main

import (
	"os"
	"time"

	"github.com/rollbar/rollbar-go"
)

func initRollbar() {
	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	rollbar.SetEnvironment("production")                 // defaults to "development"
	rollbar.SetCodeVersion("v2")                         // optional Git hash/branch/tag (required for GitHub integration)
	rollbar.SetServerHost("web.1")                       // optional override; defaults to hostname
	rollbar.SetServerRoot("github.com/heroku/myproject") // path of project (required for GitHub integration and non-project stacktrace collapsing)

	rollbar.Info("Message body goes here")
	rollbar.WrapAndWait(doSomething)
	rollbar.Wait()
	rollbar.Critical("Message body goes here")
	rollbar.Wait()
	rollbar.Error(rollbar.ERR, "Message body goes here")
	rollbar.Wait()
	rollbar.Warning("Message body goes here")
	rollbar.Wait()
	rollbar.Debug("Message body goes here")
	rollbar.Wait()
	rollbar.Wait()
	rollbar.Log(rollbar.INFO, "Message body goes here")
	rollbar.Wait()

	// call rollbar.Close() before the application exits to flush error message queue
	rollbar.Close()
}

func doSomething() {
	var timer *time.Timer = nil
	timer.Reset(10) // this will panic
}
