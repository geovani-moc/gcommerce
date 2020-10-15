package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/util"
)

//StockTemplateVariables variables in the page
type StockTemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
}

//Stock generate page
func Stock(w http.ResponseWriter, r *http.Request, root *util.Root) {

	variables := StockTemplateVariables{
		Title:       "Gerenciar estoque",
		Pages:       root.Pages,
		CurrentPage: "stock",
	}

	err := root.Templates.ExecuteTemplate(w, "stock", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
