package util

import (
	"html/template"

	"github.com/geovani-moc/gcommerce/i18n"
)

//Root base da aplicação
type Root struct {
	Templates       *template.Template
	Port            string
	Pages           []string
	SiteName        string
	Dictionaries    []i18n.Dictionary
	CurrentLanguage int // 1- english, 2- portugues
}
