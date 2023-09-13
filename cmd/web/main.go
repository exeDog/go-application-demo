package main

import (
	"fmt"
	"github.com/exedog/go-application-demo/pkg/config"
	"github.com/exedog/go-application-demo/pkg/handlers"
	"github.com/exedog/go-application-demo/pkg/render"
	"log"
	"net/http"
)

const PORT = ":9000"

func main() {
	var appConfig config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("Error creating template cache :", err)
	}

	appConfig.TemplateCache = tc
	appConfig.UseCache = true
	render.NewTemplateCache(&appConfig)

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	fmt.Println("Server running on port 9000")

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe()
	log.Fatalln(err)
}
