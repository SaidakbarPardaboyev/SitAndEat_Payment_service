package main

import (
	"fmt"
	"log"
	"net"
	"payments/config"
	pb "payments/genproto/payment"
	"payments/service"
	"payments/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().PAYMENT_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}

	paymentservice := service.NewPaymentServiceRepo(db)
	server := grpc.NewServer()
	pb.RegisterPaymentServer(server, paymentservice)

	fmt.Printf("Server is listening on port %s...", config.Load().PAYMENT_SERVICE)
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
