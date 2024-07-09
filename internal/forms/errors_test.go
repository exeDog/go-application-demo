package forms

import "testing"

func TestErrors_Add(t *testing.T) {
	e := errors{}

	e.Add("somefield", "some message")

	if e.Get("somefield") != "some message" {
		t.Error("error message not found")
	}
}

func TestErrors_Get(t *testing.T) {
	e := errors{}

	if e.Get("somefield") != "" {
		t.Error("error message found")
	}
}
