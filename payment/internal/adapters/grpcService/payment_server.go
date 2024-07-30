package grpcservice

import (
	context "context"

	"github.com/google/uuid"
	"github.com/tanasinp/go-microservices-ecommerce/payment/internal/core"
	protoPayment "github.com/tanasinp/go-microservices-ecommerce/proto/payment"
)

type paymentServiceServer struct {
	service core.PaymentService
	protoPayment.UnimplementedPaymentServiceServer
}

func NewPaymentServiceServer(service core.PaymentService) protoPayment.PaymentServiceServer {
	return &paymentServiceServer{service: service}
}

// func (paymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

func (s *paymentServiceServer) CreatePayment(ctx context.Context, req *protoPayment.CreatePaymentRequest) (*protoPayment.CreatePaymentResponse, error) {
	payment := &core.Payment{
		ID:      generateUniqueID(),
		OrderID: req.OrderId,
		UserID:  req.UserId,
		Status:  "Pending",
		Total:   req.Total,
	}
	if err := s.service.CreatePayment(payment); err != nil {
		return nil, err
	}
	res := protoPayment.CreatePaymentResponse{
		PaymentId: payment.ID,
		Status:    payment.Status,
	}
	return &res, nil
}

func (s *paymentServiceServer) GetPaymentStatus(ctx context.Context, req *protoPayment.GetPaymentStatusRequest) (*protoPayment.GetPaymentStatusResponse, error) {
	payment, err := s.service.GetPaymentStatusByID(req.PaymentId)
	if err != nil {
		return nil, err
	}
	return &protoPayment.GetPaymentStatusResponse{
		PaymentId: payment.ID,
		OrderId:   payment.OrderID,
		UserId:    payment.UserID,
		Status:    payment.Status,
		Total:     payment.Total,
	}, nil
}

func generateUniqueID() string {
	return "payment-" + uuid.New().String()
}
