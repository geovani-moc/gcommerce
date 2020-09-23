package entity

import "time"

//Client e-commerce
type Client struct {
	Code      int64
	CPF       string
	Came      string
	BirthDate time.Time
	Sexo      string //int
}
