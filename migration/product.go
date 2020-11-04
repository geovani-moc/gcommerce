package migration

import (
	"database/sql"
	"log"
	"math/rand"

	"github.com/geovani-moc/gcommerce/database"
	"github.com/geovani-moc/gcommerce/entity"
)

//CreateAllTables if not create
func CreateAllTables() error {

	return nil
}

//PopulateAllTables insert fakes itens in tables
func PopulateAllTables(quantity int64) {

}

//CreateSchemaProduct create the table in database
func CreateSchemaProduct(db *sql.DB) {
	schema := `
	create table if not exists product(
		id serial primary key,
		code integer unique,
		name text unique,
		description text,
		price real,
		quantity_stock decimal,
		status integer,
		category integer,
		id_brand_product integer
	);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

// PopulateProduct fakes products
func PopulateProduct(size int) {

	var products []entity.Product
	var product entity.Product

	columns := []string{
		"name",
		"description",
		"price",
	}

	for i := 0; i < size; i++ {
		product.Name = RandStringRunes(4)
		product.Description = RandStringRunes(4)
		product.Price = (rand.Float64() * 100) + 1

		append(products, product)
	}

	database.InsertProducts(products, columns)
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
