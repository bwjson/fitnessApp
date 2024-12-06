package service

import (
	"github.com/bwjson/fitnessApp/internal/repository"
	"github.com/bwjson/fitnessApp/pkg/auth"
	"github.com/bwjson/fitnessApp/pkg/hash"
	"github.com/bwjson/fitnessApp/pkg/logger"
	"time"
)

type Service struct {
	mongoRepo       *repository.MongoRepository
	log             logger.Logger
	tokenManager    auth.TokenManager
	hasher          hash.SHA1Hasher
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewService(mongoRepo *repository.MongoRepository, log logger.Logger, tokenManager *auth.TokenManager, hasher *hash.SHA1Hasher, accessTokenTTL, refreshTokenTTL time.Duration) *Service {
	return &Service{
		mongoRepo:       mongoRepo,
		log:             log,
		tokenManager:    *tokenManager,
		hasher:          *hasher,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}
