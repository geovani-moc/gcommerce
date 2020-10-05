package entity

import (
	"time"
)

//Sale define
type Sale struct {
	ID            int64     `json:"id"`
	Code          int64     `json:"code"`
	BeginDate     time.Time `json:"begin_date"`
	EndDate       time.Time `json:"end_date"`
	FormOfPayment int64     `json:"form_of_payment"`
	IDcart        Cart      `json:"id_cart"`
}
