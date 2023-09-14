package session

import (
	"github.com/alexedwards/scs/v2"
	"github.com/exedog/go-application-demo/pkg/config"
	"net/http"
	"time"
)

func CreateSession(config *config.AppConfig) {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = config.Production

	config.Session = session
}
