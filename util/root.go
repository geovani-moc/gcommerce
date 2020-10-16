package util

import "html/template"

//Root base da aplicação
type Root struct {
	Templates *template.Template
	Port      string
	Pages     []string
	SiteName  string
	//colocar slide de strutura de idiomas aqui
	CurrentLanguage int // 1- english, 2- portugues
}
