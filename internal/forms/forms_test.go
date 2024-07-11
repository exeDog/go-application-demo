package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "123")
	postedData.Add("b", "123")
	postedData.Add("c", "123")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows required fields missing when they are not")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	has := form.Has("whatever")

	if has {
		t.Error("Form post has fields as true when it is not")
	}

	postedData := url.Values{}
	postedData.Add("whatever", "123")

	form = New(postedData)
	has = form.Has("whatever")

	if !has {
		t.Error("Form post does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("foo", "123")

	form := New(postedData)
	form.MinLength("foo", 3)

	if !form.Valid() {
		t.Error("expected valid, got invalid")
	}

	form.MinLength("foo", 4)

	if form.Valid() {
		t.Error("expected invalid, got valid")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "foobarnotemail.com")

	form := New(postedData)
	form.IsEmail("email")

	if form.Valid() {
		t.Error("expected invalid, got valid")
	}

	postedData = url.Values{}
	postedData.Add("email", "foobar@isemail.com")

	form = New(postedData)
	form.IsEmail("email")

	if !form.Valid() {
		t.Error("expected valid, got invalid")
	}
}
