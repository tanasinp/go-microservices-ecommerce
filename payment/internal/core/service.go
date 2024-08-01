package core

import (
	adaptersOrder "github.com/tanasinp/go-microservices-ecommerce/payment/internal/adapters"
)

// primary port

type PaymentService interface {
	CreatePayment(payment *Payment) error
	GetPaymentStatusByID(paymentID string) (*Payment, error)
	UpdatePaymentStatusByID(paymentID string, status string) error
}

type paymentServiceImpl struct {
	repo         PaymentRepository
	orderService adaptersOrder.OrderService
}

func NewPaymentService(repo PaymentRepository, orderService adaptersOrder.OrderService) PaymentService {
	return &paymentServiceImpl{repo: repo, orderService: orderService}
}

func (s *paymentServiceImpl) CreatePayment(payment *Payment) error {
	if err := s.repo.SavePayment(payment); err != nil {
		return err
	}
	return nil
}

func (s *paymentServiceImpl) GetPaymentStatusByID(paymentID string) (*Payment, error) {
	payment, err := s.repo.FindPaymentStatusByID(paymentID)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *paymentServiceImpl) UpdatePaymentStatusByID(paymentID string, status string) error {
	if err := s.repo.UpdatePaymentStatusByID(paymentID, status); err != nil {
		return err
	}
	payment, err := s.GetPaymentStatusByID(paymentID) //Fetch payment details to get OrderID
	if err != nil {
		return err
	}
	if err := s.orderService.UpdateOrderStatus(payment.OrderID, status); err != nil {
		return err
	}
	return nil
}
