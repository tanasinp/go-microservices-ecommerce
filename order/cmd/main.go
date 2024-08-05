package main

import (
	"fmt"
	"log"
	"net"

	adaptersPayment "github.com/tanasinp/go-microservices-ecommerce/order/internal/adapters"
	adapters "github.com/tanasinp/go-microservices-ecommerce/order/internal/adapters/db"
	"github.com/tanasinp/go-microservices-ecommerce/order/internal/adapters/grpcService"
	"github.com/tanasinp/go-microservices-ecommerce/order/internal/core"
	protoOrder "github.com/tanasinp/go-microservices-ecommerce/proto/order"
	protoPayment "github.com/tanasinp/go-microservices-ecommerce/proto/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := setupDatabase()
	paymentService, paymentConn := setupPaymentService()
	defer paymentConn.Close()
	startGRPCServer(db, paymentService)
}

// set up the database connection and migration
func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5433", "myuser", "mypassword", "orders_db")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&core.Order{}, &core.OrderItem{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	return db
}

// set up the gRPC connection to the payment service
func setupPaymentService() (adaptersPayment.PaymentService, *grpc.ClientConn) {
	creds := insecure.NewCredentials()
	paymentConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	paymentClient := protoPayment.NewPaymentServiceClient(paymentConn)
	return adaptersPayment.NewPaymentSevice(paymentClient), paymentConn
}

// start the gRPC server for the order service
func startGRPCServer(db *gorm.DB, paymentService adaptersPayment.PaymentService) {
	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo, paymentService)
	orderServer := grpcService.NewOrderServiceServer(orderService)

	grpcServer := grpc.NewServer()
	protoOrder.RegisterOrderServiceServer(grpcServer, orderServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gRPC Order server listening on port 50051")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
