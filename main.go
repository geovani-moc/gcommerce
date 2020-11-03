package main

import (
	"fmt"
	"log"
	"os"

	"github.com/geovani-moc/gcommerce/actions"
)

/* LEMBRAR-------------------------------------------------
tambem comecar os testes
----------------------------------------------------------*/

func main() {
	arguments := os.Args

	if len(arguments) > 1 {
		execute(arguments[1:])
	}

	app := actions.BuildApp() // adicionar sub dominios par usar os idiomas
	err := app.Run()          // verificar problema de varios idiomas rodando ao mesmo tempo

	if err != nil {
		log.Fatal(err)
	}
}

func execute(commands []string) {
	for _, command := range commands {
		fmt.Println(command)
	}
}
