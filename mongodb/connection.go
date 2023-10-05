package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

func Connect(MONGODB_URI string) *mongo.Client {
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	} else {
		if os.Getenv("DEBUG") == "true" {
			println("MongoDB connection established")
		}
	}
	return MongoDBClient
}

func GetConnection() *mongo.Client {
	return MongoDBClient
}
