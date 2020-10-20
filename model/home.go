package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/util"
)

//HomeTemplateVariables variables
type HomeTemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
}

//Home initial page model
func Home(w http.ResponseWriter, r *http.Request, root *util.Root) {

	variables := HomeTemplateVariables{
		Title:       "Lista de páginas",
		Pages:       root.NamePages,
		CurrentPage: "home",
	}

	err := root.Templates.ExecuteTemplate(w, "home", variables)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
