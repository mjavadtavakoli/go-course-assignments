package service

import (
	"oms/internal/domain"
	"oms/internal/ports"
)

type OrderService struct {
	repo ports.OrderRepository
}

func NewOrderService(r ports.OrderRepository) *OrderService {
	return &OrderService{
		repo: r,
	}
}

func (s *OrderService) CreateOrder(order *domain.Order) error {
	order.Statuses = []string{"CREATED"}

	return s.repo.Create(order)
}

func (s *OrderService) GetOrder(id int64) (*domain.Order, error) {
	return s.repo.GetByID(id)
}
