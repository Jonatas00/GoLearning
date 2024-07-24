package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, template string, dados any) {
	templates.ExecuteTemplate(w, template, dados)
}

func Teste() (int, error) {
	return 0, nil
}
