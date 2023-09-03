package akademy

import (
	"embed"
	"html/template"
)

type Config struct {
	Port          uint
	Location      string
	DBName        string
	InProduction  bool
	TemplateCache map[string]*template.Template
}

//go:embed templates static
var FS embed.FS
