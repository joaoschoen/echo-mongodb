package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

func Connect() *mongo.Client {
	MONGODB_URI := os.Getenv("MONGODB_URI")
	DATABASE := os.Getenv("DATABASE")
	if MONGODB_URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	if DATABASE == "" {
		log.Fatal("You must set your 'DATABASE' environment variable.")
	}
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
