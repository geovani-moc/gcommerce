package entity

import (
	"time"
)

//Cart defines
type Cart struct {
	ID        int64     `json:"id"`
	Code      int64     `json:"code"`
	BeginDate time.Time `json:"begin_date"`
	EndDate   time.Time `json:"end_date"`
}
