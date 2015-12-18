package web

import (
	"io"

	"html/template"
)

type (
	Template struct {
		templates *template.Template
	}
)

func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
