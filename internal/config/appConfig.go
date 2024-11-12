package config

import (
	"log"
	"text/template"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Test          string
	InProduction  bool
	// Session       *scs.SessionManager
	AUTH_USER     string
	AUTH_PASSWORD string
}

func NewApp(cache bool, state bool, test string) *AppConfig {
	return &AppConfig{
		UseCache:     cache,
		Test:         test,
		InProduction: state,
	}

}
