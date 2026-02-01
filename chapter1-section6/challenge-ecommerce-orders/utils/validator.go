package utils

import (
    "errors"
    "github.com/mjavadtavakoli/challenge-ecommerce-orders/models"
)

func ValidateOrder(o models.Order) error {
	if o.Email == "" {
		return errors.New("email is required")
	}
	if o.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}
