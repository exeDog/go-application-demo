package main

import (
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var mh myHandler

	h := Nosurf(&mh)

	switch h.(type) {
	case http.Handler:
		// do nothing; test passed
	default:
		t.Error("type is not http.Handler")

	}
}

func TestLoadSession(t *testing.T) {
	var mh myHandler

	h := LoadSession(&mh)

	switch h.(type) {
	case http.Handler:
		// do nothing; test passed
	default:
		t.Error("type is not http.Handler")

	}
}
