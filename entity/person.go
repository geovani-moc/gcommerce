package entity

import "time"

//Person e-commerce
type Person struct {
	Code          int64
	CPF           string
	Came          string
	BirthDate     time.Time
	Type          int64 // admin, client or official
	Sexo          int64
	authenticated bool
	product       Product //olhar se faz sentido remover
	cart          Cart
	report        Report
	contact       Contact
	address       Address
}
