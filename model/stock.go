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
func Stock(w http.ResponseWriter, r *http.Request) {

	variables := StockTemplateVariables{
		Title: "Gerenciar estoque",
	}

	pageTemplate, err := template.ParseFiles("view/stock.html", "view/base/footer.html",
		"view/base/header.html", "view/base/side-bar.html")

	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	err = pageTemplate.ExecuteTemplate(w, "stock", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
