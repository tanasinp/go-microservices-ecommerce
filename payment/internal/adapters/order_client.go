package adapters

import (
	"context"

	protoOrder "github.com/tanasinp/go-microservices-ecommerce/proto/order"
)

type OrderService interface {
	UpdateOrderStatus(orderID string, status string) error
}

type orderService struct {
	orderClient protoOrder.OrderServiceClient
}

func NewOrderService(orderClient protoOrder.OrderServiceClient) OrderService {
	return orderService{orderClient: orderClient}
}

func (base orderService) UpdateOrderStatus(orderID string, status string) error {
	req := protoOrder.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  status,
	}
	res, err := base.orderClient.UpdateOrderStatus(context.Background(), &req)
	if err != nil {
		return err
	}
	_ = res
	return nil
}
