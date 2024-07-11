package helpers

import (
	"fmt"
	"github.com/exedog/go-application-demo/internal/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(rw http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(rw, http.StatusText(status), status)
}

func ServerError(rw http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Printf("Server error with error of %s, stackTrace: %s", err.Error(), trace)
	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
