package internal

import (
	"context"
	"github.com/bwjson/fitnessApp/internal/models"
)

type MongoRepository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
}
