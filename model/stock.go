package model

import (
	"log"
	"net/http"
	"strconv"

	"github.com/geovani-moc/gcommerce/entity"
	"github.com/geovani-moc/gcommerce/util"
)

//StockTemplateVariables variables in the page
type StockTemplateVariables struct {
	Title              string
	Pages              []string
	CurrentPage        string
	ErrorInsertProduct string
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

//StockInsertProduct for method post
func StockInsertProduct(w http.ResponseWriter, r *http.Request, root *util.Root) {
	if r.FormValue("cust_id") == "insert-product" {
		variables := StockTemplateVariables{
			Title:       "Gerenciar estoque",
			Pages:       root.Pages,
			CurrentPage: "stock",
		}

		product := entity.Product{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
		}
		var err error

		product.Code, err = strconv.ParseInt(r.FormValue("code"), 10, 0)
		if err != nil {
			log.Print(err)
		}

		product.Price, err = strconv.ParseFloat(r.FormValue("price"), 10)
		if err != nil {
			log.Print(err)
		}

		log.Print(product)

		err = root.Templates.ExecuteTemplate(w, "stock", variables)
		if nil != err {
			log.Print("Erro ao gerar template insert product, ", err)
		}
	}
}
