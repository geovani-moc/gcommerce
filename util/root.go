package util

import (
	"html/template"

	"github.com/geovani-moc/gcommerce/i18n"
)

//Root base da aplicação
type Root struct {
	Templates       *template.Template
	Port            string
	NamePages       []string
	SiteName        string
	Dictionaries    map[string]i18n.Dictionary
	CurrentLanguage string
	Dictionary      i18n.Dictionary
}
