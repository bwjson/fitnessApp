package internal

import (
	"context"
	"github.com/bwjson/fitnessApp/internal/models"
)

type MongoRepository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
	GetUser(ctx context.Context, email, password string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}
