package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/database"
	"github.com/geovani-moc/gcommerce/util"

	"github.com/geovani-moc/gcommerce/entity"
)

// ProductTemplateVariables variables for template parser
type ProductTemplateVariables struct {
	Title       string
	Products    []entity.Product
	Pages       []string
	CurrentPage string
	URL         string
}

// Product controller
func Product(w http.ResponseWriter, r *http.Request, root *util.Root) {
	var err error
	variables := ProductTemplateVariables{
		Title:       "Todos produtos",
		Pages:       root.NamePages,
		CurrentPage: "product",
		URL:         root.URL,
	}

	variables.Products, err = database.GetAllProducts()
	if nil != err {
		log.Println("Erro ao carregar dados(produtos) do banco de dados, ", err)
	}

	err = root.Templates.ExecuteTemplate(w, "product", variables)

	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
