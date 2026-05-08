package domain

import "time"

type Order struct {
	ID           int64     `json:"id"`
	CustomerName string    `json:"customer_name"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
