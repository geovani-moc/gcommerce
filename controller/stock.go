package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/util"

	"github.com/geovani-moc/gcommerce/model"
)

//Stock redirect to page stock
func Stock(w http.ResponseWriter, r *http.Request, root *util.Root) {

	if r.Method == http.MethodPost {
		model.StockInsertProduct(w, r, root)
		return
	}

	model.Stock(w, r, root)
}
