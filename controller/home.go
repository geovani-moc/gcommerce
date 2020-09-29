package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
)

//Home initial page controller
func Home(w http.ResponseWriter, r *http.Request) {
	model.Home(w, r)
}
