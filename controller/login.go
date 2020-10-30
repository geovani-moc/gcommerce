package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
	"github.com/geovani-moc/gcommerce/util"
)

//Login redirect
func Login(w http.ResponseWriter, r *http.Request, root *util.Root) {
	model.Login(w, r, root)
}
