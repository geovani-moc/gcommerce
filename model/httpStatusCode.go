package model

import (
	"html/template"
	"log"
	"net/http"
)

//Code404TemplateVariables variaveis da pagina
type Code404TemplateVariables struct {
	Title string
}

//Code404 generate page
func Code404(w http.ResponseWriter, r *http.Request, globalTemplate *template.Template) {
	variables := Code404TemplateVariables{
		Title: "Página não encontrada",
	}

	err := globalTemplate.ExecuteTemplate(w, "not-found", variables)
	if nil != err {
		log.Println("Error ao gerar pagina nao encontrada, ", err)
	}
}
