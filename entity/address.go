package entity

//Address define
type Address struct {
	ID               int64  `json:"id"`
	Code             int64  `json:"code"`
	PublicPlace      int64  `json:"public_place"`
	NamePublicPlace  string `json:"name_public_place"`
	City             int64  `json:"city"`
	State            int64  `json:"state"`
	Country          int64  `json:"country"`
	PostalCode       string `json:"postal_code"`
	Complement       string `json:"complement"`
	Number           int64  `json:"number"`
	ComplementNumber string `json:"complement_number"`
}
