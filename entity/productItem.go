package entity

//ProductItem defines
type ProductItem struct {
	ID        int64   `json:"id"`
	Code      int64   `json:"code"`
	Quantity  float64 `json:"quantity"`
	IDproduct Product `json:"id_product"`
}
