package entity

// Product define
type Product struct {
	code          int64
	name          string
	description   string
	price         float64
	quantityStock float64
	status        int64
	brand         Brand
}
