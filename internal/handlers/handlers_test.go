package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	Key   string
	Value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
	params             []postData
}{
	{"home", "/", "GET", http.StatusOK, []postData{}},
	{"about", "/about", "GET", http.StatusOK, []postData{}},
	{"contact", "/contact", "GET", http.StatusOK, []postData{}},
	{"gs", "/generals-quarters", "GET", http.StatusOK, []postData{}},
	{"ms", "/majors-suite", "GET", http.StatusOK, []postData{}},
	{"sa", "/search-availability", "GET", http.StatusOK, []postData{}},
	{"mr", "/make-reservation", "GET", http.StatusOK, []postData{}},
	{"gs", "/generals-quarters", "GET", http.StatusOK, []postData{}},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	// Create a new server using the "routes" just created
	// and then start the server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	// Loop through the tests
	for _, e := range theTests {
		// Create a new request
		req, _ := http.NewRequest(e.method, ts.URL+e.url, nil)

		// Add any needed form data
		if len(e.params) > 0 {
			q := req.URL.Query()
			for _, x := range e.params {
				q.Add(x.Key, x.Value)
			}

			req.URL.RawQuery = q.Encode()
		}

		// Make the request
		res, err := ts.Client().Do(req)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		// Check the status code
		if res.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
		}
	}
}
