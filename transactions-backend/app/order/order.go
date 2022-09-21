package order

type Product struct {
	ID          string  `json:"_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	V           int     `json:"__v"`
}

type Order struct {
	Document_id    string  `json:"document_id"`
	Client_name    string  `json:"client_name"`
	Client_email   string  `json:"client_email"`
	Total_price    float64 `json:"total_price"`
	Status         string  `json:"status"`
	Products       string  `json:"products"`
	Created_at     string  `json:"created_at"`
	V              int     `json:"__v"`
	Transaction_id string  `json:"transaction_id"`
}

func NewOrder() *Order {
	return &Order{}
}
