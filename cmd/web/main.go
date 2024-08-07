package main

import (
	"encoding/gob"
	"fmt"
	"github.com/exedog/go-application-demo/internal/config"
	"github.com/exedog/go-application-demo/internal/handlers"
	"github.com/exedog/go-application-demo/internal/helpers"
	"github.com/exedog/go-application-demo/internal/models"
	"github.com/exedog/go-application-demo/internal/render"
	"github.com/exedog/go-application-demo/internal/session"
	"log"
	"net/http"
	"os"
)

const PORT = "localhost:9009"

var appConfig config.AppConfig

func main() {
	gob.Register(models.Reservation{})
	appConfig.InProduction = false
	appConfig.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session.CreateSession(&appConfig)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("Error creating template cache :", err)
	}

	appConfig.TemplateCache = tc
	appConfig.UseCache = false
	render.NewTemplates(&appConfig)

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&appConfig)

	fmt.Println("Server running on port 9000")

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe()
	log.Fatalln(err)
}
