package entity

// Product define
type Product struct {
	id            int64
	Code          int64
	Name          string
	Description   string
	Price         float64
	QuantityStock float64
	Status        int64
	Cold          float64
	Category      int64
	BrandProduct  Brand
}
