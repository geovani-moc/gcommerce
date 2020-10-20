package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
	"github.com/geovani-moc/gcommerce/util"
)

//Profile redirect to generator page
func Profile(w http.ResponseWriter, r *http.Request, root *util.Root) {
	if r.FormValue("edit") == "true" {
		model.ProfileEdit(w, r, root)
	}

	model.Profile(w, r, root)
}
