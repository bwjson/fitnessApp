package server

import (
	"github.com/bwjson/fitnessApp/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	//log     logger.Logger
	cfg     *config.Config
	mongoDB *mongo.Client
}
