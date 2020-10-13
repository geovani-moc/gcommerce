package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/entity"
	"github.com/geovani-moc/gcommerce/util"
)

//ProfileVariablesTemplate definition
type ProfileVariablesTemplate struct {
	Title   string
	profile entity.Person
}

//Profile generate page
func Profile(w http.ResponseWriter, r *http.Request, root *util.Root) {

	variable := ProfileVariablesTemplate{
		Title: "Perfil",
	}

	err := root.Templates.ExecuteTemplate(w, "profile", variable)
	if nil != err {
		log.Print("Erro ao fazer o parse da pagina de perfil, ", err)
	}
}
