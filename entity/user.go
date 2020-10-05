package entity

//User for the login
type User struct {
	ID       int64  `json:"id"`
	Code     int64  `json:"code"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Status   int64  `json:"status"`
	Type     int64  `json:"type"`
	IDperson Person `json:"id_person"`
}
