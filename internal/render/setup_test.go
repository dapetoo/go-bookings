package render

import (
	"github.com/alexedwards/scs/v2"
	"github.com/dapetoo/go-bookings/internal/config"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	// change this to true when in production
	app.InProduction = true

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}
