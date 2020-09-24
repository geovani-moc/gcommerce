package entity

import (
	"time"
)

//Cart defines
type Cart struct {
	Code      int64
	BeginDate time.Time
	EndDate   time.Time
}
