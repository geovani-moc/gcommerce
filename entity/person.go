package entity

import "time"

//Person e-commerce
type Person struct {
	Code      int64
	CPF       string
	Came      string
	BirthDate time.Time
	Type      int64
	Sexo      int64 //int
	product   Product
	cart      Cart
	report    Report
	contact   Contact
	address   Address
}
