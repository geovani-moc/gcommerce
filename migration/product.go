package migration

import (
	"database/sql"
	"log"
	"math/rand"

	"github.com/geovani-moc/gcommerce/database"
	"github.com/geovani-moc/gcommerce/entity"
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

// PopulateProduct fakes products
func PopulateProduct(size int) {

	var product entity.Product
	for i := 0; i < size; i++ {
		product.Name = RandStringRunes(4)
		product.Description = RandStringRunes(4)
		product.Price = (rand.Float64() * 100) + 1
	}

	database.InsertProduct(product)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandStringRunes string aleatoria
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
