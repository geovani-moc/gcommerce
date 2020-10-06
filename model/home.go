package model

import (
	"html/template"
	"log"
	"net/http"
)

//HomeTemplateVariables variables
type HomeTemplateVariables struct {
	Title string
	Pages []string
}

//Home initial page model
func Home(w http.ResponseWriter, r *http.Request, globalTemplate *template.Template) {

	variables := HomeTemplateVariables{
		Title: "Lista de páginas",
		Pages: []string{
			"product",
			"stock",
		},
	}

	err := globalTemplate.ExecuteTemplate(w, "home", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
