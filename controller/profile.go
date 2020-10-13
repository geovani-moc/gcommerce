package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
	"github.com/geovani-moc/gcommerce/util"
)

//Profile redirect to generator page
func Profile(w http.ResponseWriter, r *http.Request, root *util.Root) {
	model.Profile(w, r, root)
}
