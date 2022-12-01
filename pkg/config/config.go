package config

import (
	"log"
	"text/template"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
