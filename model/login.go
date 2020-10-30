package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/util"
)

//LoginTemplateVariables variables for generate
type LoginTemplateVariables struct {
	Title       string
	Pages       []string
	CurrentPage string
}

//Login generate page login
func Login(w http.ResponseWriter, r *http.Request, root *util.Root) {

	variables := LoginTemplateVariables{
		Title:       "Autenticação",
		Pages:       root.NamePages,
		CurrentPage: "login",
	}

	err := root.Templates.ExecuteTemplate(w, "login", variables)

	if nil != err {
		log.Print("Erro ao gerar pagina de login, ", err)
	}
}
