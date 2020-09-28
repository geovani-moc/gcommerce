package entity

// Product define
type Product struct {
	Code          int64
	Name          string
	Description   string
	Price         float64
	QuantityStock float64
	Status        int64
	BrandProduct  Brand
}
