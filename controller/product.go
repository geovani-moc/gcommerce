package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
	"github.com/geovani-moc/gcommerce/util"
)

//Product redirect
func Product(w http.ResponseWriter, r *http.Request, root *util.Root) {

	model.Product(w, r, root)
}
