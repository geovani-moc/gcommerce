package model

import (
	"html/template"
	"log"
	"net/http"
)

//StockTemplateVariables variables in the page
type StockTemplateVariables struct {
	Title string
}

//Stock generate page
func Stock(w http.ResponseWriter, r *http.Request, globalTemplate *template.Template) {

	variables := StockTemplateVariables{
		Title: "Gerenciar estoque",
	}

	err := globalTemplate.ExecuteTemplate(w, "stock", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
