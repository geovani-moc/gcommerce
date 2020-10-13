package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
	"github.com/geovani-moc/gcommerce/util"
)

//Code404 redirect to page not foud
func Code404(w http.ResponseWriter, r *http.Request, root *util.Root) {
	model.Code404(w, r, root)
}
