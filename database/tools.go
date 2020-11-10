package database

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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
