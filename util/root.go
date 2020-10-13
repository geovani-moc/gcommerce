package util

import "html/template"

//Root base da aplicação
type Root struct {
	Templates *template.Template
	Port      string
	Pages     []string
}
