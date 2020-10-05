package entity

// Brand defines
type Brand struct {
	ID          int64  `json:"id"`
	Code        int64  `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}
