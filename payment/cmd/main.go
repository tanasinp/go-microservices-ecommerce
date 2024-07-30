package main

import (
	"fmt"
	"log"
	"net"

	adapters "github.com/tanasinp/go-microservices-ecommerce/payment/internal/adapters/db"
	grpcService "github.com/tanasinp/go-microservices-ecommerce/payment/internal/adapters/grpcService"
	"github.com/tanasinp/go-microservices-ecommerce/payment/internal/core"
	protoPayment "github.com/tanasinp/go-microservices-ecommerce/proto/payment"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection setup
	host := "localhost"
	port := "5434"
	user := "myuser"
	password := "mypassword"
	dbname := "payments_db"
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&core.Payment{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	// Setup services and gRPC server
	paymentRepo := adapters.NewGormPaymentRepository(db)
	paymentService := core.NewPaymentService(paymentRepo)
	paymentServer := grpcService.NewPaymentServiceServer(paymentService)

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	// Register the payment service with the gRPC server
	protoPayment.RegisterPaymentServiceServer(grpcServer, paymentServer)

	// Start the gRPC server
	fmt.Println("gRPC Payment server listening on port 50052")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
