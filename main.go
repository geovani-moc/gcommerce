package main

import (
	"log"

	"github.com/geovani-moc/gcommerce/actions"
)

func main() {
	app := actions.BuildApp()
	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
