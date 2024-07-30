package adapters

import (
	"context"

	proto "github.com/tanasinp/go-microservices-ecommerce/proto/payment"
)

type PaymentService interface {
	CreatePayment(orderID string, userID string, total float64) error
}

type paymentService struct {
	paymentClient proto.PaymentServiceClient
}

func NewPaymentSevice(paymentClient proto.PaymentServiceClient) PaymentService {
	return paymentService{paymentClient: paymentClient}
}

func (base paymentService) CreatePayment(orderID string, userID string, total float64) error {
	req := proto.CreatePaymentRequest{
		OrderId: orderID,
		UserId:  userID,
		Total:   total,
	}
	res, err := base.paymentClient.CreatePayment(context.Background(), &req)
	if err != nil {
		return err
	}
	_ = res
	return nil
}
