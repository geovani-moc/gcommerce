package migration

import (
	"database/sql"
	"log"
)

//CreateSchemaProduct create the table in database
func CreateSchemaProduct(db *sql.DB) {
	schema := `
	create table if not exists product(
		id serial primary key,
		code integer unique,
		name text unique,
		description text,
		price real,
		quantity_stock ,
		status integer real,
		category integer real,
		id_brand_product integer
	);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}
