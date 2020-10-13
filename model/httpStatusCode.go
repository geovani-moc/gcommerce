package model

import (
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/util"
)

//Code404TemplateVariables variaveis da pagina
type Code404TemplateVariables struct {
	Title string
}

//Code404 generate page
func Code404(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variables := Code404TemplateVariables{
		Title: "Página não encontrada",
	}

	err := root.Templates.ExecuteTemplate(w, "not-found", variables)
	if nil != err {
		log.Println("Error ao gerar pagina nao encontrada, ", err)
	}
}
