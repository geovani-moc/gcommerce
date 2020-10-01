package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/geovani-moc/gcommerce/database"
)

//reponse format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//GetAllProducts api rest return all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//get all the products in the db
	//products, err := database.GetAllProducts()
	products, err := database.GetFakeProducts()

	if err != nil {
		log.Fatal("Unable to get all products. \n", err)
	}

	//send all the products as response
	json.NewEncoder(w).Encode(products)
}
