package main

import (
	"html/template"
	"log"

	"github.com/geovani-moc/gcommerce/actions"
)

//TemplateApp temporary local
var TemplateApp *template.Template

func main() {
	app := actions.BuildApp()
	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
