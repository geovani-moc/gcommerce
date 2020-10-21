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
		Pages:       root.NamePages,
		Person:      database.GetFakePerson(),
		CurrentPage: "profile",
	}

	err := root.Templates.ExecuteTemplate(w, "profile", variable)
	if nil != err {
		log.Print("Erro ao fazer o parse da pagina de perfil, ", err)
	}
}

//ProfileEdit for editions
func ProfileEdit(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variable := ProfileVariablesTemplate{
		Title:       "Perfil",
		Pages:       root.NamePages,
		Person:      database.GetFakePerson(),
		CurrentPage: "profile",
	}

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

	err := root.Templates.ExecuteTemplate(w, "profile", variable)
	if nil != err {
		log.Print("Erro ao gerar templateedit profile, ", err)
	}
}

//ProfileEditLanguage modify language
func ProfileEditLanguage(w http.ResponseWriter, r *http.Request, root *util.Root) {
	variable := ProfileVariablesTemplate{
		Title:       "Perfil",
		Pages:       root.NamePages,
		Person:      database.GetFakePerson(),
		CurrentPage: "profile",
	}

	root.CurrentLanguage = r.FormValue("language")
	// update language from root
	err := root.Templates.ExecuteTemplate(w, "profile", variable)
	if nil != err {
		log.Print("Erro ao gerar templateedit profile, ", err)
	}
}
