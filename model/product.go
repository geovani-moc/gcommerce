package model

import (
	"html/template"
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/database"

	"github.com/geovani-moc/gcommerce/entity"
)

// ProductTemplateVariables variables for template parser
type ProductTemplateVariables struct {
	Title    string
	Products []entity.Product
}

// Product controller
func Product(w http.ResponseWriter, r *http.Request, globalTemplate *template.Template) {
	var err error
	variables := ProductTemplateVariables{
		Title: "Todos produtos",
	}

	variables.Products, err = database.GetFakeProducts()
	if nil != err {
		log.Println("Erro ao carregar dados(produtos) do banco de dados, ", err)
	}

	err = globalTemplate.ExecuteTemplate(w, "product", variables)

	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
