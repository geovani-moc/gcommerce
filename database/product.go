package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/geovani-moc/gcommerce/entity"
)

//InsertProduct insert one product in database
func InsertProduct(product entity.Product) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `
	insert into products (id, code, name, description, 
		price, quantity_stock, status, category) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	returning id`

	var id int64

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if nil != err {
		log.Fatal("erro ao iniciar a transação, ", err)
	}

	// _, err = tx.ExecContext(ctx, sqlStatement, product.ID, product.Code,
	// 	product.Name, product.Description, product.Price,
	// 	product.QuantityStock, product.Status, product.Category)
	// if nil != err {
	// 	tx.Rollback()
	// 	log.Fatal(err)
	// }

	row := tx.QueryRowContext(ctx, sqlStatement, product.ID, product.Code,
		product.Name, product.Description, product.Price,
		product.QuantityStock, product.Status, product.Category)

	if nil != row {
		tx.Rollback()
		log.Fatal("Erro ao realizar a transação, ")
	}
	err = tx.Commit()
	if nil != err {
		log.Fatal("erro ao realizar o commit, ", err)
	}

	//retornando id invalido
	return id
}

//InsertProducts insert various products in database with one SQL
func InsertProducts(products []entity.Product, columns []string) []int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := "insert into products ("

	sizeColumns := len(columns)
	for i := 0; i < sizeColumns-1; i++ {
		sqlStatement = sqlStatement + columns[i] + ", "
	}
	sqlStatement = sqlStatement + columns[sizeColumns-1] + ") values ("

	for i := 0; i < len(products); i++ {
		for j := 0; j < sizeColumns; j++ {
			value := (i * sizeColumns) + j
			sqlStatement = sqlStatement + "$" + strconv.FormatInt(int64(value), 10)
		}
	} //funcao incomlpeta
	return nil
}

//GetAllProducts return prodducts in database
func GetAllProducts() ([]entity.Product, error) {
	db := createConnection()
	defer db.Close()

	var products []entity.Product

	sqlStatement := `SELECT * FROM  product`

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(
			&product.ID, &product.Code, &product.Name,
			&product.Description, &product.Price, &product.QuantityStock,
			&product.Status, &product.Category)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		products = append(products, product)
	}
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

	err := row.Scan(
		&product.ID, &product.Code, &product.Name,
		&product.Description, &product.Price, &product.QuantityStock,
		&product.Status, &product.Category)

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

	sqlStatement := `update product set 
	code=$2, name=$3, description=$4, price=$5, quantity_stock=$6,
	status=$7, category=$8 
	where id=$1`

	response, err := db.Exec(sqlStatement,
		id, product.Code, product.Name, product.Description,
		product.Price, product.QuantityStock, product.Status,
		product.Category)

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
