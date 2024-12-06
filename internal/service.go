package internal

import (
	"context"
	"github.com/bwjson/fitnessApp/internal/dto"
	"github.com/bwjson/fitnessApp/internal/models"
)

type Service interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, input dto.LoginInput) (dto.TokenResponse, error)
	GetProfileInfo(ctx context.Context, email string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}
