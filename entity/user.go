package entity

//User for the login
type User struct {
	Code   int64
	Login  string
	Senha  string
	Status int
	Type   string //int
	Email  string
}
