package entity

import "time"

//Person e-commerce
type Person struct {
	ID            int64     `json:"id"`
	Code          int64     `json:"code"`
	CPF           string    `json:"cpf"`
	Name          string    `json:"name"`
	BirthDate     time.Time `json:"birth_date"`
	Type          int64     `json:"type"` // admin, client or official
	Sex           int64     `json:"sex"`
	Authenticated bool      `json:"authenticated"`
	IDproduct     Product   `json:"id_product"` //olhar se faz sentido remover
	IDcart        Cart      `json:"id_cart"`
	IDreport      Report    `json:"id_report"`
	IDcontact     Contact   `json:"id_contact"`
	IDaddress     Address   `json:"id_address"`
}
