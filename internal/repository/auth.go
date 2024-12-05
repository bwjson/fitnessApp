package repository

import (
	"context"
	"fmt"
	"github.com/bwjson/fitnessApp/internal/models"
	"github.com/bwjson/fitnessApp/pkg/http_errors"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	FitnessDB       = "fitnessApp"
	UsersCollection = "users"
)

func (r *MongoRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	collection := r.mongoDB.Database(FitnessDB).Collection(UsersCollection)

	user.RegisteredAt = time.Now().UTC()
	user.LastVisitAt = time.Now().UTC()

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeErr := range writeException.WriteErrors {
				if writeErr.Code == 11000 {
					return nil, err
				}
			}
		}
		return nil, errors.Wrap(err, "InsertOne")
	}

	fmt.Printf("InsertedID: %v\n", result.InsertedID)

	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.Wrap(err, "objectId")
	}

	user.UserID = objectId

	return user, nil
}

func (r *MongoRepository) GetUser(ctx context.Context, email, password string) (*models.User, error) {
	collection := r.mongoDB.Database(FitnessDB).Collection(UsersCollection)

	filter := bson.D{{
		"$and", bson.A{
			bson.D{{"email", email}},
			bson.D{{"password", password}},
		},
	}}

	var user models.User

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http_errors.UserNotFound
		}
		return nil, errors.Wrap(err, "FindOne")
	}

	return &user, nil
}
