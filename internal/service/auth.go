package service

import (
	"context"
	"github.com/bwjson/fitnessApp/internal/models"
)

func (s *Service) Create(ctx context.Context, user *models.User) (*models.User, error) {
	return s.mongoRepo.Create(ctx, user)
}
