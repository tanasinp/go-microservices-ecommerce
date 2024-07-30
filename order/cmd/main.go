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
	host := "localhost"
	port := "5433"
	user := "myuser"
	password := "mypassword"
	dbname := "orders_db"
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connec database")
	}
	db.AutoMigrate(&core.Order{}, &core.OrderItem{})

	creds := insecure.NewCredentials()
	paymentConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer paymentConn.Close()
	paymentClient := protoPayment.NewPaymentServiceClient(paymentConn)
	paymentService := adaptersPayment.NewPaymentSevice(paymentClient)

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo, paymentService) //
	orderServer := grpcService.NewOrderServiceServer(orderService)

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	protoOrder.RegisterOrderServiceServer(grpcServer, orderServer)

	fmt.Println("gRPC Order server listening on port 50051")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
