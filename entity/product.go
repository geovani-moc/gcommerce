package entity

// Product define
type Product struct {
	ID            int64   `json:"id"`
	Code          int64   `json:"code"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	QuantityStock float64 `json:"quantity_stock"`
	Status        int64   `json:"status"`
	//Cold          float64 `json:"cold"`
	Category int64 `json:"category"`
	IDbrand  Brand `json:"id_brand_product"`
}
