package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	x := New(url.Values{})
	if x == nil {
		t.Error("New function returned something empty")

	}

}

func TestForm_Required(t *testing.T) {
	// Initialize the form with some data
	data := url.Values{}
	data.Add("name", "John")
	data.Add("email", "")

	// Create a new Form object
	form := New(data)

	// Test the Required function
	form.Required("name", "email")

	// Check if there is an error for the "email" field
	if _, exists := form.Errors["email"]; !exists {
		t.Error("Expected an error for missing 'email', but didn't get one")
	}

	// Check if there is no error for the "name" field
	if _, exists := form.Errors["name"]; exists {
		t.Error("Did not expect an error for 'name', but got one")
	}
}

func TestHas(t *testing.T) {
	//creating a new request
	r, _ := http.NewRequest("POST", "/anyurl", nil)
	//creating a form with empty values
	form := New(url.Values{})
	//below line must return true becausethe form creating is empty and we are posting that and then checking with Has function
	has := form.Has("anyfieldLikeNameOrEmail", r)
	if has {
		t.Error("Form Shows Has Fields When some or all fields are empty")
	}
	//checking after posting value whether it is giving false means we posted value but it is giving that we haveot posted any value

	postedData := New(url.Values{})
	postedData.Add("anyfield", "rjrjrjrj")
	has = postedData.Has("anyfield", r)
	if !has {
		t.Error("Form shows field is blank but we have posted a value")
	}

}
func TestMinLength(t *testing.T) {
	r, _ := http.NewRequest("POST", "/anyurl", nil)
	form := New(url.Values{})
	form.Add("anyfield", "w")
	min := form.MinLength("anyfield", 2, r)
	if min {
		t.Error("The field is having the length less than needed but MinLength fn validates to having >= required length")
	}
}
func TestISEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/anyurl", nil)
	form := New(r.PostForm) //postsempty form
	form.ISEmail("X")
	if form.Valid() {
		t.Error("Form shows valid email for non existing field")
	}
	//for posted values valid email
	postedValues := url.Values{}
	postedValues.Add("email", "rohan3296@gmail.com")
	form = New(postedValues)
	form.ISEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when entered a valid one!!!")
	}

}
