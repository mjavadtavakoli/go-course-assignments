package ports

import "oms/internal/core/domain"

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id int64) (*domain.Order, error)
}
