package services

import (
	"errors"
	"time"

	"oms/internal/core/domain"
	"oms/internal/core/ports"
)

type OrderService struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(customerName string, status string) (*domain.Order, error) {
	if customerName == "" {
		return nil, errors.New("customer_name is required")
	}
	if status == "" {
		status = "PENDING"
	}

	order := &domain.Order{
		CustomerName: customerName,
		Status:       status,
		CreatedAt:    time.Now(),
	}

	err := s.repo.Create(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrderByID(id int64) (*domain.Order, error) {
	if id <= 0 {
		return nil, errors.New("invalid order id")
	}

	return s.repo.GetByID(id)
}
