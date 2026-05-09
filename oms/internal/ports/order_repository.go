package ports

import "oms/internal/domain"

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id int64) (*domain.Order, error)
	GetAll() ([]domain.Order, error)
	Update(order *domain.Order) error
}
