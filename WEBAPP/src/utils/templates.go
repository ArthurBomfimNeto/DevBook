package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Carregar templates insere os templates html na variavel templates
func CarrregarTemplates() {
	// templates recebe todos os arquivos com extenção .html
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
