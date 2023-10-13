package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOne(collection string, query primitive.D) *mongo.SingleResult {
	// GET DB CONNECTION
	db := GetConnection()
	DATABASE := os.Getenv("DATABASE")
	coll := db.Database(DATABASE).Collection(collection)

	return coll.FindOne(
		context.TODO(),
		query,
	)
}
func InsertOne(collection string, data primitive.D) (*mongo.InsertOneResult, error) {
	// GET DB CONNECTION
	db := GetConnection()
	DATABASE := os.Getenv("DATABASE")
	coll := db.Database(DATABASE).Collection(collection)

	return coll.InsertOne(
		context.TODO(),
		data,
	)
}
