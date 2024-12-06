package service

import (
	"context"
	"github.com/bwjson/fitnessApp/internal/dto"
	"github.com/bwjson/fitnessApp/internal/models"
)

func (s *Service) Create(ctx context.Context, user *models.User) (*models.User, error) {
	user.Password = s.hasher.GenerateHashedPassword(user.Password)
	return s.mongoRepo.Create(ctx, user)
}

func (s *Service) Login(ctx context.Context, input dto.LoginInput) (dto.TokenResponse, error) {
	tokenResponse := dto.TokenResponse{}
	input.Password = s.hasher.GenerateHashedPassword(input.Password)

	_, err := s.mongoRepo.GetUser(ctx, input.Email, input.Password)

	if err != nil {
		return tokenResponse, err
	}

	// access token
	accessToken, err := s.tokenManager.AccessTokenGen(input, s.accessTokenTTL)
	if err != nil {
		return tokenResponse, err
	}
	tokenResponse.AccessToken = accessToken

	// refresh token
	refreshToken, err := s.tokenManager.RefreshTokenGen(s.refreshTokenTTL)
	if err != nil {
		return tokenResponse, err
	}
	tokenResponse.RefreshToken = refreshToken

	return tokenResponse, nil
}

func (s *Service) GetProfileInfo(ctx context.Context, email string) (*models.User, error) {
	return s.mongoRepo.GetUserByEmail(ctx, email)
}

func (s *Service) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.mongoRepo.GetAllUsers(ctx)
}
