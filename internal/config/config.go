package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	LogInfo       *log.Logger
	Production    bool
	Session       *scs.SessionManager
}
