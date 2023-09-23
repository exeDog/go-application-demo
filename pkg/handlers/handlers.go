package handlers

import (
	"github.com/exedog/go-application-demo/pkg/config"
	"github.com/exedog/go-application-demo/pkg/models"
	"github.com/exedog/go-application-demo/pkg/render"
	"net/http"
)

var Repo *Repository

type (
	Repository struct {
		App *config.AppConfig
	}
)

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "about.page.html", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "majors.page.html", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "make-reservation.page.html", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "contact.page.html", &models.TemplateData{})
}
