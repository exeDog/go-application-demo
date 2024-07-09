package render

import (
	"github.com/exedog/go-application-demo/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	td := models.TemplateData{}

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value in the template data is nil")
	}

	if result == nil {
		t.Error("result should be of type *models.TemplateData")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww mockResponse

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})

	if err != nil {
		t.Error("error should be nil")
	}
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"

	_, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
