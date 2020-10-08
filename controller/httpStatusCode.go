package controller

import (
	"html/template"
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
)

//Code404 redirect to page not foud
func Code404(w http.ResponseWriter, r *http.Request, globalTemplate *template.Template) {
	model.Code404(w, r, globalTemplate)
}
