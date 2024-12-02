package service

import (
	"github.com/bwjson/fitnessApp/internal/repository"
	"github.com/bwjson/fitnessApp/pkg/logger"
)

type Service struct {
	mongoRepo *repository.MongoRepository
	log       logger.Logger
}

func NewService(mongoRepo *repository.MongoRepository, log logger.Logger) *Service {
	return &Service{
		mongoRepo: mongoRepo,
		log:       log,
	}
}
