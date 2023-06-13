package forms

import (
	"net/http/httptest"
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

func TestForm_Has(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever", r)
	if has {
		t.Error("Form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "b")
	form = New(postedData)

	has = form.Has("a", r)
	if !has {
		t.Error("Shows Form does not shows when it should")
	}
}
