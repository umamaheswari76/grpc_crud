package main

import (
	"context"
	"crud_basic/config"

	// "crud_basic/constants"
	"crud_basic/services"
	"fmt"
	"net"

	pro "crud_basic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	mongoClient := config.ConnectDatabase()
	services.CustomerCollection = config.GetCollection("crud", "customer", mongoClient)
	services.Ctx = context.Background()

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	
	s := grpc.NewServer()
	reflection.Register(s)
	pro.RegisterCustomerServiceServer(s, &services.Server{})

	fmt.Println("Server listening on ", ":50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
