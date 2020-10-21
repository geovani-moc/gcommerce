package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
	"github.com/geovani-moc/gcommerce/util"
)

//Profile redirect to generator page
func Profile(w http.ResponseWriter, r *http.Request, root *util.Root) {
	option := r.FormValue("option")

	switch option {
	case "edit_profile":
		model.ProfileEdit(w, r, root)

	case "edit_language":
		model.ProfileLanguage(w, r, root)
	}

	model.Profile(w, r, root)
}
