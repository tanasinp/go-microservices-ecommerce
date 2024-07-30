package grpcService

//grpc communication
import (
	context "context"

	"github.com/google/uuid"
	"github.com/tanasinp/go-microservices-ecommerce/order/internal/core"
	protoOrder "github.com/tanasinp/go-microservices-ecommerce/proto/order"
)

type orderServiceServer struct {
	service core.OrderService
	protoOrder.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(service core.OrderService) protoOrder.OrderServiceServer {
	return &orderServiceServer{service: service}
}

// func (orderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

func (s *orderServiceServer) CreateOrder(ctx context.Context, req *protoOrder.CreateOrderRequest) (*protoOrder.CreateOrderResponse, error) {
	orderID := generateOrderUniqueID()
	order := &core.Order{
		ID:      orderID,
		UserID:  req.UserId,
		Items:   convertItems(req.Items),
		Address: req.Address,
		Total:   req.Total,
		Status:  "Pending",
	}
	for i := range order.Items {
		order.Items[i].OrderID = order.ID
	}
	if err := s.service.CreateOrder(order); err != nil {
		return nil, err
	}
	res := protoOrder.CreateOrderResponse{
		OrderId: order.ID,
	}
	return &res, nil
}

func (s *orderServiceServer) GetOrder(ctx context.Context, req *protoOrder.GetOrderRequest) (*protoOrder.GetOrderResponse, error) {
	order, err := s.service.GetOrder(req.OrderId)
	if err != nil {
		return nil, err
	}
	res := protoOrder.GetOrderResponse{
		OrderId: order.ID,
		UserId:  order.UserID,
		Items:   convertOrderItems(order.Items),
		Address: order.Address,
		Total:   order.Total,
		Status:  order.Status,
	}
	return &res, nil
}

func convertItems(items []*protoOrder.OrderItem) []core.OrderItem {
	var orderItems []core.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, core.OrderItem{
			ID:        generateOrderItemUniqueID(),
			ProductID: item.ProductId,
			Quantity:  int(item.Quantity),
			Price:     item.Price,
			OrderID:   "",
		})
	}
	return orderItems
}

func convertOrderItems(items []core.OrderItem) []*protoOrder.OrderItem {
	var pbItems []*protoOrder.OrderItem
	for _, item := range items {
		pbItems = append(pbItems, &protoOrder.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
		})
	}
	return pbItems
}

func generateOrderUniqueID() string {
	return "order-" + uuid.New().String()
}
func generateOrderItemUniqueID() string {
	return "orderItem-" + uuid.New().String()
}
