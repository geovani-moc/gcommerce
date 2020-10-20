package main

import (
	"log"

	"github.com/geovani-moc/gcommerce/actions"
)

func main() {
	app := actions.BuildApp() // adicionar sub dominios par usar os idiomas
	err := app.Run()          // verificar problema de varios idiomas rodando ao mesmo tempo

	if err != nil {
		log.Fatal(err)
	}
}
