package forms

import (
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	// Create a url.Values object
	data := url.Values{"name": []string{"John Doe"}, "age": []string{"30"}}

	// Create a new Form object
	form := New(data)

	// Check that the Form object has no errors
	if len(form.Errors) != 0 {
		t.Errorf("Form.Errors should be empty")
	}
}
