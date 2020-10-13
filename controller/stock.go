package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/geovani-moc/gcommerce/util"

	"github.com/geovani-moc/gcommerce/entity"
	"github.com/geovani-moc/gcommerce/model"
)

//Stock redirect to page stock
func Stock(w http.ResponseWriter, r *http.Request, root *util.Root) {

	if r.Method == http.MethodPost {
		if r.FormValue("cust_id") == "insert-product" {
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
		}
	}

	model.Stock(w, r, root)
}
