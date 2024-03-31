package render

import (
	"bytes"
	"fmt"
	"github.com/exedog/go-application-demo/pkg/config"
	"github.com/exedog/go-application-demo/pkg/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var appConfig *config.AppConfig

func NewTemplateCache(a *config.AppConfig) {
	appConfig = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func ShowTemplate(w http.ResponseWriter, templ string, td *models.TemplateData, r *http.Request) {
	var tc map[string]*template.Template

	if appConfig.UseCache {
		tc = appConfig.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
		appConfig.TemplateCache = tc
	}

	t, ok := tc[templ]
	if !ok {
		log.Fatalln("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatalln("Error executing template :", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatalln("Error writing template to browser :", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	fmt.Println("Creating template cache...")

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}

	return cache, nil
}
