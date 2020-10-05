package entity

import (
	"time"
)

//Report defines
type Report struct {
	ID         int64     `json:"id"`
	Code       int64     `json:"code"`
	CreateDate time.Time `json:"createDate"`
	Name       string    `json:"name"`
	Category   int64     `json:"category"`
	IDsale     Sale      `json:"id_sale"`
}
