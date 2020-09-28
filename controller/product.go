package controller

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/model"
)

//Product redirect
func Product(w http.ResponseWriter, r *http.Request) {
	// var err error
	// app.TemplateApp, err = template.ParseGlob("/view/*.html")

	// if err != nil {
	// 	panic(err)
	// }

	model.Product(w, r)
}
