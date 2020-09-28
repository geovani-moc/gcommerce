package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/entity"
)

// ProductTemplateVariables variables for template parser
type ProductTemplateVariables struct {
	Title    string
	Products []entity.Product
}

// Product controller
func Product(w http.ResponseWriter, r *http.Request) {

	variables := ProductTemplateVariables{
		Title:    "Todos produtos",
		Products: fakeProducts(),
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

func fakeProducts() []entity.Product {
	products := make([]entity.Product, 0)

	products = append(products, entity.Product{
		Code:          1,
		Name:          "Notebook lenovo",
		Description:   "Notebook super rapido com i5 e placa de video dedicada!!!",
		Price:         10.2,
		QuantityStock: 2,
		Status:        1,
	})

	return products
}
