package model

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
func Product(w http.ResponseWriter, r *http.Request, globalTemplate *template.Template) {

	variables := ProductTemplateVariables{
		Title:    "Todos produtos",
		Products: fakeProducts(),
	}

	err := globalTemplate.ExecuteTemplate(w, "product", variables)

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
		Price:         5000.00,
		QuantityStock: 2,
		Status:        1,
	})

	products = append(products, entity.Product{
		Code:          1,
		Name:          "Notebook lenovo",
		Description:   "Notebook super rapido com i9 e placa de video dedicada!!!",
		Price:         3010.20,
		QuantityStock: 2,
		Status:        1,
	})

	products = append(products, entity.Product{
		Code:          1,
		Name:          "Notebook Asus",
		Description:   "Notebook  com i3",
		Price:         2010.20,
		QuantityStock: 223,
		Status:        1,
	})

	products = append(products, entity.Product{
		Code:          1,
		Name:          "Notebook tectois",
		Description:   "Notebook amd e placa de video dedicada.",
		Price:         3100.00,
		QuantityStock: 434,
		Status:        1,
	})

	return products
}
