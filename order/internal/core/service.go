package core

import (
	adaptersPayment "github.com/tanasinp/go-microservices-ecommerce/order/internal/adapters"
)

// primary port
type OrderService interface {
	CreateOrder(order *Order) error
	GetOrder(id string) (*Order, error)
}

// business logic
type orderServiceImpl struct {
	repo           OrderRepository
	paymentService adaptersPayment.PaymentService
}

func NewOrderService(repo OrderRepository, paymentService adaptersPayment.PaymentService) OrderService {
	return &orderServiceImpl{repo: repo, paymentService: paymentService}
}

func (s *orderServiceImpl) CreateOrder(order *Order) error {
	if err := s.repo.Save(order); err != nil {
		return err
	}
	if err := s.paymentService.CreatePayment(order.ID, order.UserID, order.Total); err != nil {
		return err
	}
	return nil
}

func (s *orderServiceImpl) GetOrder(id string) (*Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
