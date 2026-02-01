package models

type Order struct {
	ID     string  `json:"id"`
	Email  string  `json:"email"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

