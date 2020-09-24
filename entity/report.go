package entity

import (
	"time"
)

//Report defines
type Report struct {
	code       int64
	createDate time.Time
	name       string
	category   int64
	sale       Sale
}
