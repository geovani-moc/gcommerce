package controller

import (
	"html/template"
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
)

//Home initial page controller
func Home(w http.ResponseWriter, r *http.Request, globalTemplate *template.Template) {
	model.Home(w, r, globalTemplate)
}
