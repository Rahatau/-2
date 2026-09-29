package config

import "html/template"

type AppConfig struct {
	UseCache       bool
	
	TemplatesCache map[string]*template.Template
}