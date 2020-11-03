package main

import (
	"log"
	"os"

	"github.com/geovani-moc/gcommerce/migration"

	"github.com/geovani-moc/gcommerce/actions"
)

/* LEMBRAR-------------------------------------------------
tambem comecar os testes
----------------------------------------------------------*/

func main() {
	arguments := os.Args

	if len(arguments) > 1 {
		executeCommand(arguments[1:])
	}

	app := actions.BuildApp()
	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func executeCommand(commands []string) {
	for _, command := range commands {
		switch command {
		case "-m":
			log.Println("Iniciando crição de tabelas ...")
			err := migration.CreateAllTables()
			if nil != err {
				log.Println("Erro na criação de tabelas, ", err)
			} else {
				log.Println("Iniciando população de tabelas ...")
				migration.PopulateAllTables(1000)
			}
		default:
			log.Println("O comando \"", command, "\" não foi reconhecido.")
		}

	}
}
