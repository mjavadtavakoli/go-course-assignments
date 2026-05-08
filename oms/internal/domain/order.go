package domain

type Order struct {
	ID           int64    `json:"id"`
	CustomerName string   `json:"customer_name"`
	Amount       float64  `json:"amount"`
	Statuses     []string `json:"statuses"`
}
