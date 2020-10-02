package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/geovani-moc/gcommerce/entity"

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

//CreateProduct api get product
func CreateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var product entity.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Print("Unable to decode the request body, ", err)
	}

	insertID := database.InsertProduct(product)
	returnResponse := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(returnResponse)
}

//GetProduct return one product by id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r) //mudar para _app
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Print("Unable to convert the string into int ", err)
	}

	product, err := database.GetProduct(int64(id))
	if err != nil {
		log.Print("Unable to get product. ", err)
	}

	json.NewEncoder(w).Encode(product)
}

//UpdateProduct update product's data
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Print("Unable to convert the string into int. ", err)
	}

	var product entity.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Print("Unable to decode the request body. ", err)
	}

	updatedRows := database.UpdateProduct(int64(id), product)
	message := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	returnResponse := response{
		ID:      int64(id),
		Message: message,
	}

	json.NewEncoder(w).Encode(returnResponse)
}

//DeleteProduct remove produc by id
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	// corrigir com(convercao para int e depois int 64):
	// id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Print("Unable to convert the string into int. ", err)
	}

	deletedRows := database.DeleteProduct(int64(id))
	message := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	returnResponse := response{
		ID:      int64(id),
		Message: message,
	}

	json.NewEncoder(w).Encode(returnResponse)
}
