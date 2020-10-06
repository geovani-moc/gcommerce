package controller

import (
	"html/template"
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
)

//Product redirect
func Product(w http.ResponseWriter, r *http.Request, globalTemplate *template.Template) {

	model.Product(w, r, globalTemplate)
}
