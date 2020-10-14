package database

import (
	"time"

	"github.com/geovani-moc/gcommerce/entity"
)

//GetFakePerson return a fake person
func GetFakePerson() entity.Person {
	return entity.Person{
		ID:            1,
		Code:          2030,
		CPF:           "111-222-333-90",
		Name:          "Abcd Efghik",
		BirthDate:     time.Now(),
		Type:          2,
		Sex:           1,
		Authenticated: true,
	}
}
