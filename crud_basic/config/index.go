package config

import (
	"context"
	"crud_basic/constants"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase ()(*mongo.Client){
	ctx,_ := context.WithTimeout(context.Background(),100*time.Second)
	mongoClient := options.Client().ApplyURI(constants.ConnectionString)
	mongoConn, err := mongo.Connect(ctx, mongoClient)
	if err != nil {
		log.Println("Can't connect to mongo Client: ",err)
	}
	return mongoConn
}

func GetCollection(db string,coll string ,client *mongo.Client)(collection *mongo.Collection){
	mongoCollection := client.Database(db).Collection(coll)
	return mongoCollection
}
