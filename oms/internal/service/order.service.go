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

	logger.Log.Printf(
		"processing create order: customer=%s",
		order.CustomerName,
	)

	order.Statuses = []string{"CREATED"}

	err := s.repo.Create(order)

	if err != nil {

		logger.Log.Printf(
			"create order failed in service: customer=%s error=%v",
			order.CustomerName,
			err,
		)

		return err
	}

	logger.Log.Printf(
		"create order completed: id=%d",
		order.ID,
	)

	return nil
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
