package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Quando for colocado em produção mudar
//essa informações para um arquivo .env
const (
	host         = "localhost"
	port         = 5432
	user         = "postgres"
	password     = "postgres"
	databaseName = "gcommerce"
)

//create connection with postgres db
func createConnection() *sql.DB {
	postgresURL := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, databaseName)

	db, err := sql.Open("postgres", postgresURL)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected.")

	return db
}
