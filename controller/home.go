package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
	"github.com/geovani-moc/gcommerce/util"
)

//Home initial page controller
func Home(w http.ResponseWriter, r *http.Request, root *util.Root) {
	model.Home(w, r, root)
}
