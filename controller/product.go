package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
)

//Product redirect
func Product(w http.ResponseWriter, r *http.Request) {

	model.Product(w, r)
}
