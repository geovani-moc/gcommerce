package database

import (
	"database/sql"
	"encoding/json"
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

//CreateInsert create the sentence sql to insert into table
func CreateInsert(quantity int, table string, columns []string) string {
	sqlStatement := "insert into " + table + " ("

	sizeColumns := len(columns)
	for i := 0; i < sizeColumns-1; i++ {
		sqlStatement = sqlStatement + columns[i] + ", "
	}
	sqlStatement = sqlStatement + columns[sizeColumns-1] + ") values"

	for i := 0; i < quantity; i++ {
		sqlStatement += "("
		for j := 0; j < sizeColumns; j++ {
			value := (i * sizeColumns) + j + 1
			sqlStatement = sqlStatement + "$" + strconv.FormatInt(int64(value), 10) + ","
		}
		sqlStatement = sqlStatement[:len(sqlStatement)-1] + "),"
	}

	sqlStatement = sqlStatement[:len(sqlStatement)-1]
	return sqlStatement
}

func structToMap(in *[]interface{}) *[]map[string]interface{} {
	var out []map[string]interface{}
	var temp map[string]interface{}

	for _, row := range *in {
		inrec, _ := json.Marshal(row)
		json.Unmarshal(inrec, &temp)
		out = append(out, temp)
	}

	return &out
}

//SQLvalues extract values from entity for setence sql
func SQLvalues(data *[]map[string]interface{}, columns []string) *[]interface{} {
	var values []interface{}

	for _, row := range *data {
		for i := 0; i < len(columns); i++ {
			cell := row[columns[i]]
			values = append(values, fmt.Sprintf("%v", cell))
		}
	}

	return &values
}
