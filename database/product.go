package database

import (
	"errors"
	"log"

	"github.com/geovani-moc/gcommerce/entity"
)

//GetAllProducts return prodducts in database
func GetAllProducts() ([]entity.Product, error) {
	//create the postgres db connection
	db := createConnection()

	//close the postgres db connection
	defer db.Close()

	var products []entity.Product

	//create the select sql query
	sqlStatement := `SELECT * FROM users`

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
			Code: 1,
		},
		{
			Code: 2,
		},
		{
			Code: 3,
		},
	}

	return products, nil
}
