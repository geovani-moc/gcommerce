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
