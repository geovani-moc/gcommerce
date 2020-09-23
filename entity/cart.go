package entity

import (
	"time"
)

//Cart defines
type Cart struct {
	code      int64
	beginDate time.Time
	endDate   time.Time
}
