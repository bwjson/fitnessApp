package service

import (
	"context"
	"github.com/bwjson/fitnessApp/internal/dto"
	"github.com/bwjson/fitnessApp/internal/models"
	"github.com/bwjson/fitnessApp/pkg/auth"
	"github.com/bwjson/fitnessApp/pkg/hash"
	"time"
)

const (
	salt       = "3n92rn329fjids9"
	signingKey = "nf3823848h3290fb289"
	ttl        = time.Minute * 30
)

func (s *Service) Create(ctx context.Context, user *models.User) (*models.User, error) {
	hashes := hash.NewSHA1Hasher(salt)
	user.Password = hashes.GenerateHashedPassword(user.Password)
	return s.mongoRepo.Create(ctx, user)
}

func (s *Service) Login(ctx context.Context, input dto.LoginInput) (dto.TokenResponse, error) {
	tokenResponse := dto.TokenResponse{}
	hashes := hash.NewSHA1Hasher(salt)
	input.Password = hashes.GenerateHashedPassword(input.Password)

	_, err := s.mongoRepo.GetUser(ctx, input.Email, input.Password)

	if err != nil {
		return tokenResponse, err
	}

	manager, err := auth.NewTokenManager(signingKey, ttl)
	if err != nil {
		return tokenResponse, err
	}

	// access token
	accessToken, err := manager.AccessTokenGen(input)
	if err != nil {
		return tokenResponse, err
	}
	tokenResponse.AccessToken = accessToken

	// refresh token
	refreshToken, err := manager.RefreshTokenGen()
	if err != nil {
		return tokenResponse, err
	}
	tokenResponse.RefreshToken = refreshToken

	return tokenResponse, nil
}
