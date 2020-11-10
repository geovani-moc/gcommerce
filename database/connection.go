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

//Quando for colocado em produção mudar
//essa informações para um arquivo .env
const (
	host         = "localhost"
	port         = 5432
	user         = "postgres"
	password     = "postgres"
	databaseName = "gcommerce"
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

	fmt.Println("Successfully connected.")

	return db
}

//CreateInsert create the sentence sql to insert into table
func CreateInsert(quantity int, table string, columns []string) string {
	sqlStatement := "insert into " + table + " ("

	sizeColumns := len(columns)
	for i := 0; i < sizeColumns-1; i++ {
		sqlStatement = sqlStatement + columns[i] + ", "
	}
	sqlStatement = sqlStatement + columns[sizeColumns-1] + ") values("

	for i := 0; i < quantity; i++ {
		for j := 0; j < sizeColumns; j++ {
			value := (i * sizeColumns) + j + 1
			sqlStatement = sqlStatement + "$" + strconv.FormatInt(int64(value), 10) + ","
		}
	}

	sqlStatement = sqlStatement[:len(sqlStatement)-1] + ")"

	return sqlStatement
}
