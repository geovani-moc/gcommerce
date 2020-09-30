package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
)

//Stock redirect to page stock
func Stock(w http.ResponseWriter, r *http.Request) {
	model.Stock(w, r)
}
