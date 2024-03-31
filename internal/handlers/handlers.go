package handlers

import (
	"encoding/json"
	"github.com/exedog/go-application-demo/internal/config"
	"github.com/exedog/go-application-demo/internal/models"
	"github.com/exedog/go-application-demo/internal/render"
	"net/http"
)

var Repo *Repository

type (
	Repository struct {
		App *config.AppConfig
	}

	jsonResponse struct {
		OK      bool   `json:"ok"`
		Message string `json:"message"`
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
	render.ShowTemplate(w, "home.page.html", &models.TemplateData{}, r)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "about.page.html", &models.TemplateData{}, r)
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "generals.page.html", &models.TemplateData{}, r)
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "majors.page.html", &models.TemplateData{}, r)
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "make-reservation.page.html", &models.TemplateData{}, r)
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "search-availability.page.html", &models.TemplateData{}, r)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.ShowTemplate(w, "contact.page.html", &models.TemplateData{}, r)
}

func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Post search availability"))
	if err != nil {
		return
	}
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
