package repository

import (
	"context"
	"fmt"
	"github.com/bwjson/fitnessApp/internal/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	fitnessDB       = "fitnessApp"
	usersCollection = "users"
)

func (r *MongoRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	collection := r.mongoDB.Database(fitnessDB).Collection(usersCollection)

	user.RegisteredAt = time.Now().UTC()
	user.LastVisitAt = time.Now().UTC()

	fmt.Printf("Before InsertOne: %+v\n", user)

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
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
