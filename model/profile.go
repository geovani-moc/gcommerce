package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/database"

	"github.com/geovani-moc/gcommerce/entity"
	"github.com/geovani-moc/gcommerce/util"
)

//ProfileVariablesTemplate definition
type ProfileVariablesTemplate struct {
	Title       string
	Pages       []string
	Person      entity.Person
	CurrentPage string
}

//Profile generate page
func Profile(w http.ResponseWriter, r *http.Request, root *util.Root) {

	variable := ProfileVariablesTemplate{
		Title:       "Perfil",
		Pages:       root.Pages,
		Person:      database.GetFakePerson(),
		CurrentPage: "profile",
	}

	err := root.Templates.ExecuteTemplate(w, "profile", variable)
	if nil != err {
		log.Print("Erro ao fazer o parse da pagina de perfil, ", err)
	}
}
