package entity

import (
	"time"
)

//Sale define
type Sale struct {
	code          int64
	beginDate     time.Time
	endDate       time.Time
	formOfPayment int64
	cart          Cart
}
