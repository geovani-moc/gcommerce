package controller

import (
	"html/template"
	"log"
	"net/http"
)

// ProductTemplateVariables variables for template parser
type ProductTemplateVariables struct {
	Title string
}

// Product controller
func Product(w http.ResponseWriter, r *http.Request) {

	variables := ProductTemplateVariables{
		Title: "Todos produtos",
	}

	pageTemplate, err := template.ParseFiles("view/product.html")
	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	err = pageTemplate.Execute(w, variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
