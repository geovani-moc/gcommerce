package migration

import (
	"database/sql"
	"math/rand"

	"github.com/geovani-moc/gcommerce/database"
	"github.com/geovani-moc/gcommerce/entity"
)

//CreateAllTables if not create
func CreateAllTables() error {
	db := database.CreateConnection()
	defer db.Close()
	CreateSchemaProduct(db)
	return nil
}

//PopulateAllTables insert fakes itens in tables
func PopulateAllTables(quantity int64) {

}

//CreateSchemaProduct create the table in database
func CreateSchemaProduct(db *sql.DB) error {
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
		return err
	}

	return nil
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
		product.Name = RandStringRunes(10)
		product.Description = RandStringRunes(20)
		product.Price = (rand.Float64() * 100) + 1
		product.Code = (rand.Int63n(999999))

		products = append(products, product)
	}

	database.InsertProducts(products, columns)
}

var letterRunes = []rune("abcd efghijkl mnopqrstuvw xyzABCDEFGHIJKL MNOPQRSTUVWXYZ ")

//RandStringRunes string aleatoria
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
