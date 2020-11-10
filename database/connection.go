package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	//driver postgresql
	_ "github.com/lib/pq"
)

//CreateConnection with postgres db
func CreateConnection() *sql.DB {

	host := os.Getenv("HOST")
	port, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 64)
	if nil != err {
		log.Println("Erro ao converter POSTGRES_PORT, ", err)
		port = 5432
	}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")

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

	return db
}
