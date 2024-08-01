package main

import (
	"fmt"
	"log"
	"net"

	adaptersOrder "github.com/tanasinp/go-microservices-ecommerce/payment/internal/adapters"
	adapters "github.com/tanasinp/go-microservices-ecommerce/payment/internal/adapters/db"
	grpcService "github.com/tanasinp/go-microservices-ecommerce/payment/internal/adapters/grpcService"
	"github.com/tanasinp/go-microservices-ecommerce/payment/internal/core"
	protoOrder "github.com/tanasinp/go-microservices-ecommerce/proto/order"
	protoPayment "github.com/tanasinp/go-microservices-ecommerce/proto/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := setupDatabase()
	orderService := setupOrderService()
	startGRPCServer(db, orderService)
}

// set up the database connection and migration
func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5434", "myuser", "mypassword", "payments_db")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&core.Payment{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	return db
}

// set up the gRPC connection to the order service
func setupOrderService() adaptersOrder.OrderService {
	creds := insecure.NewCredentials()
	orderConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer orderConn.Close()

	orderClient := protoOrder.NewOrderServiceClient(orderConn)
	return adaptersOrder.NewOrderService(orderClient)
}

// start the gRPC server for the payment service
func startGRPCServer(db *gorm.DB, orderService adaptersOrder.OrderService) {
	paymentRepo := adapters.NewGormPaymentRepository(db)
	paymentService := core.NewPaymentService(paymentRepo, orderService)
	paymentServer := grpcService.NewPaymentServiceServer(paymentService)

	grpcServer := grpc.NewServer()
	protoPayment.RegisterPaymentServiceServer(grpcServer, paymentServer)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gRPC Payment server listening on port 50052")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
