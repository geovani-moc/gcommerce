package controller

import (
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/entity"

	"github.com/geovani-moc/gcommerce/model"
	"github.com/geovani-moc/gcommerce/util"
)

//Profile redirect to generator page
func Profile(w http.ResponseWriter, r *http.Request, root *util.Root) {
	if r.FormValue("edit") == "true" {
		person := entity.Person{
			Name: r.FormValue("name"),
			CPF:  r.FormValue("cpf"),
		}
		sexValue := r.FormValue("sex")
		log.Print(sexValue)

		if sexValue == "Homen" {
			person.Sex = 1
		} else if sexValue == "Mulher" {
			person.Sex = 2
		} else {
			person.Sex = 0
		}

		log.Print("Editando: ", person)
	}

	model.Profile(w, r, root)
}
