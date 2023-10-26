package services

import (
	"context"
	"crud_basic/models"
	pro "crud_basic/proto"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerCollection *mongo.Collection
	Ctx                context.Context
)

func (s *Server) Create(ctx context.Context, req *pro.CreateCustomerRequest) (*pro.CreateCustomerResponse, error) {
	db := &models.Customer{
		Id:       req.Id,
		Name:     req.Name,
		Password: req.Password,
		Contact:  req.Contact,
	}

	res, err := CustomerCollection.InsertOne(Ctx, &db)
	if err != nil {
		log.Println("can't insert customer:  ", err)
	}

	fmt.Println("customer created, : ", res.InsertedID)
	return &pro.CreateCustomerResponse{
		Response: "Customer created Successfully",
	}, nil
}

func (s *Server) Read(ctx context.Context, req *pro.ReadCustomerRequest) (*pro.ReadCustomerResponse, error) {
	// db := &models.Customer{
	// 	Id:       req.Id,
	// }
	query := bson.M{"id": req.Id}

	var read *models.Customer
	err := CustomerCollection.FindOne(Ctx, query).Decode(&read)
	if err != nil {
		return nil, err
	}

	fmt.Println("customer read, : \n")
	return &pro.ReadCustomerResponse{
		Id:      read.Id,
		Name:    read.Name,
		Contact: read.Contact,
	}, nil
}

func (s *Server) Update(ctx context.Context, req *pro.UpdateCustomerRequest) (*pro.UpdateCustomerResponse, error) {

	filter := bson.D{{"_id", req.Id}}
	query := bson.D{{"$set", bson.D{{"contact", "9876543210"}}}}
	res, err := CustomerCollection.UpdateOne(Ctx, filter, query)
	if err != nil {
		log.Println("Can't update customer: ", err)
	}

	fmt.Println("customer updated, : \n", res.UpsertedID)
	return &pro.UpdateCustomerResponse{
		Response: "Customer updated Succesfully",
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *pro.DeleteCustomerRequest) (*pro.DeleteCustomerResponse, error) {
	filter := bson.M{"id": req.Id}
	res, err := CustomerCollection.DeleteOne(Ctx, filter)
	if err != nil {
		log.Println("Cant delete customer: ", err)
	}

	fmt.Println("customer deleted, : \n", res.DeletedCount)
	return &pro.DeleteCustomerResponse{
		Response: "Customer deleted Succesfully",
	}, err
}
