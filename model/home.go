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
func Home(w http.ResponseWriter, r *http.Request) {

	variables := HomeTemplateVariables{
		Title: "Lista de páginas",
		Pages: []string{
			"product",
			"stock",
		},
	}

	pageTemplate, err := template.ParseFiles("view/home.html", "view/base/footer.html",
		"view/base/header.html", "view/base/side-bar.html")

	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	err = pageTemplate.Execute(w, variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
