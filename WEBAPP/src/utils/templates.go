package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

//CarregarTemplates insere os templates html na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html")) // os arquivos que vão ser jogados dentro da variavel templates são todos com extensão .html
}

func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
