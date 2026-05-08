package ports

import "oms/internal/core/domain"

type OrderUseCase interface {
	CreateOrder(customerName string, status string) (*domain.Order, error)
	GetOrderByID(id int64) (*domain.Order, error)
}
