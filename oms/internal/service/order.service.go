package service

import (
	"oms/internal/domain"
	"oms/internal/logger"
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
	//logger:
	logger.Log.Println("Created Order !")

	return s.repo.Create(order)
}

func (s *OrderService) GetOrder(id int64) (*domain.Order, error) {
	logger.Log.Println("Call api Getorder !")

	return s.repo.GetByID(id)
}

func (s *OrderService) GetAllOrders() ([]domain.Order, error) {
	logger.Log.Println("Call api GetAll !")

	return s.repo.GetAll()
}

func (s *OrderService) UpdateOrder(order *domain.Order) error {
	logger.Log.Println("Updated Order !")

	return s.repo.Update(order)
}
