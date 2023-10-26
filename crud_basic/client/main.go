package main

import (
	"context"
	pro "crud_basic/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Can't dial grpc client: ", err)
	}

	client := pro.NewCustomerServiceClient(conn)

	createCustReq := &pro.CreateCustomerRequest{
		Id:       "7",
		Name:     "ggg",
		Password: "1234",
		Contact:  "123456789",
	}
	createCustRes, err := client.Create(context.Background(),createCustReq)
	if err != nil {
		log.Fatalf("Create customer err: %v", err)
	}
	log.Println(createCustRes.Response)


	readCustReq := &pro.ReadCustomerRequest{
		Id:      "7",
	}
	readCustRes, err := client.Read(context.Background(),readCustReq)
	if err != nil {
		log.Fatalf("Create customer err: %v", err)
	}
	log.Println(readCustRes.Name)

	updateCustReq := &pro.UpdateCustomerRequest{
		Id:      "7",
		Contact: "9876543210",
	}
	updateCustRes, err := client.Update(context.Background(),updateCustReq)
	if err != nil {
		log.Fatalf("Create customer err: %v", err)
	}
	log.Println(updateCustRes.Response)

	deleteCustReq := &pro.DeleteCustomerRequest{
		Id:      "7",
	}
	deleteCustRes, err := client.Delete(context.Background(),deleteCustReq)
	if err != nil {
		log.Fatalf("Create customer err: %v", err)
	}
	log.Println(deleteCustRes.Response)



}