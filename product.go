package main

//Product of the store
type Product struct {
	code        int64
	name        string
	description string
	price       float32
	status      string
	//quantidade em estoque, assim elimina  a entidade estoque
	//nome da marca
}
