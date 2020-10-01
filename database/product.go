package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/geovani-moc/gcommerce/entity"
)

//InsertProduct in database
func InsertProduct(product entity.Product) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `insert into products (code, name, description, price) VALUES ($1, $2, $3, $4) returning id`

	var id int64

	err := db.QueryRow(sqlStatement, product.Code, product.Name, product.Description, product.Price)
	if err != nil {
		log.Print("Unable to execute the query ", err)
	}
	log.Print("Inserted a single record ", id)

	return id
}

//GetAllProducts return prodducts in database
func GetAllProducts() ([]entity.Product, error) {
	//create the postgres db connection
	db := createConnection()

	//close the postgres db connection
	defer db.Close()

	var products []entity.Product

	//create the select sql query
	sqlStatement := `SELECT * FROM  product`

	//execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	//close the statement
	defer rows.Close()

	//iterate over the rows
	for rows.Next() {
		var product entity.Product

		//unmarshal the row object to user
		err = rows.Scan(&product.Code, &product.Name, &product.Description, &product.Price)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the user slice
		products = append(products, product)
	}

	//return empty product on error
	return products, errors.New("vazio")
}

// GetFakeProducts fakes products
func GetFakeProducts() ([]entity.Product, error) {
	products := []entity.Product{
		{
			Code:          1,
			Name:          "Notebook lenovo",
			Description:   "Notebook super rapido com i5 e placa de video dedicada!!!",
			Price:         5000.00,
			QuantityStock: 2,
			Status:        1,
		},
		{
			Code:          2,
			Name:          "Notebook lenovo",
			Description:   "Notebook super rapido com i9 e placa de video dedicada!!!",
			Price:         3010.20,
			QuantityStock: 2,
			Status:        1,
		},
		{
			Code:          3,
			Name:          "Notebook Asus",
			Description:   "Notebook  com i3",
			Price:         2010.20,
			QuantityStock: 223,
			Status:        2,
		},
		{
			Code:          4,
			Name:          "Notebook tectois",
			Description:   "Notebook amd e placa de video dedicada.",
			Price:         3100.00,
			QuantityStock: 434,
			Status:        1,
		},
	}

	return products, nil
}

//GetProduct return product by id
func GetProduct(id int64) (entity.Product, error) {
	db := createConnection()
	defer db.Close()

	var product entity.Product

	sqlStatement := `select * FROM products where id=$1`
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&product.Code, &product.Name, &product.Description, &product.Price)

	switch err {
	case sql.ErrNoRows:
		log.Print("No rows were retourned")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Print("Unable to scan the row ", err)
	}

	return product, err
}

//UpdateProduct refresh the product values by id
func UpdateProduct(id int64, product entity.Product) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `update product set code=$2, name=$3, description=$4, price=$5 where id=$1`
	response, err := db.Exec(sqlStatement, id, product.Code, product.Name, product.Description, product.Price)

	if err != nil {
		log.Print("Unable to execute the query ", err)
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		log.Print("Error while checking the affected rows. ", err)
	}
	log.Print("Total rows/record affected ", rowsAffected)

	return rowsAffected
}

//DeleteProduct remove product by id
func DeleteProduct(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `delete from products where id=$1`
	response, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Print("Unable to execute the query. ", err)
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		log.Print("Error while checking the affected rows. ", err)
	}

	log.Print("Total rows/record affected ", err)

	return rowsAffected
}
