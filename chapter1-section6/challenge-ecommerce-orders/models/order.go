package models

type Order struct {
	// چرا ایدی و به صورت رشته گرفتی؟
	ID     string  `json:"id"`
	Email  string  `json:"email"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

