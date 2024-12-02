package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	mongoDB *mongo.Client
}

func NewMongoRepository(mongoDB *mongo.Client) *MongoRepository {
	return &MongoRepository{mongoDB: mongoDB}
}
