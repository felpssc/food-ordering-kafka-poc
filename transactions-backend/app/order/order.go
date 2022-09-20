package main

type Product struct {
	_id         string
	name        string
	description string
	price       float64
	_v          int
}

type Order struct {
	document_id  string
	client_name  string
	client_email string
	total_price  float64
	status       string
	products     []Product
	created_at   string
	_v           int
}

func NewOrder() *Order {
	return &Order{}
}
