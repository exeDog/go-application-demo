package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Name:     "csrf_token",
		Secure:   appConfig.Production,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}

func LoadSession(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}
