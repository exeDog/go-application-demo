package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/exedog/go-application-demo/internal/config"
	"github.com/exedog/go-application-demo/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})
	testApp.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type mockResponse struct{}

func (*mockResponse) Header() http.Header {
	return http.Header{}
}

func (*mockResponse) WriteHeader(int) {}

func (*mockResponse) Write(b []byte) (int, error) {
	return len(b), nil
}
