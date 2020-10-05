package entity

//Contact define
type Contact struct {
	ID    int64  `json:"id"`
	Code  int64  `json:"code"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
