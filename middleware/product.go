package middleware

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/entity"
)

//reponse format
type response struct {
	ID      int64  `json: "id, omitempty"`
	Message string `json: "message,omitempty`
}

//GetAllProducts api rest return all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//get all the products in the db
	products, err := getAllProducts()

	if err != nil {
		log.Fatal("Unable to get all products. %v", err)
	}

	//send all the products as response
	json.NewEncoder(w).Encode(products)
}

//getAllProducts
func getAllProducts() ([]entity.Product, error) {
	// //create the postgres db connection
	// db := createConnection()

	// //close the postgres db connection
	// defer db.Close()

	var products []entity.Product

	// //create the select sql query
	// sqlStatement := `SELECT * FROM users`

	// //execute the sql statement
	// rows, err := db.Query(sqlStatement)

	// if err != nil {
	// 	log.Fatalf("Unable to execute the query. %v", err)
	// }

	// //close the statement
	// defer rows.Close()

	// //iterate over the rows
	// for rows.Next() {
	// 	var user model.User

	// 	//unmarshal the row object to user
	// 	err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	// 	if err != nil {
	// 		log.Fatalf("Unable to scan the row. %v", err)
	// 	}

	// 	// append the user in the user slice
	// 	users = append(users, user)
	// }

	//return empty product on error
	return products, errors.New("vazio")
}
