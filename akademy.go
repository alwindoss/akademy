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
	FS            embed.FS
	TemplateCache map[string]*template.Template
}
