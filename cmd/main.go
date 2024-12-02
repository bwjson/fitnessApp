package main

import (
	"context"
	"github.com/bwjson/fitnessApp/config"
	"github.com/bwjson/fitnessApp/internal/repository"
	"github.com/bwjson/fitnessApp/internal/server"
	"github.com/bwjson/fitnessApp/internal/service"
	"github.com/bwjson/fitnessApp/internal/transport/rest"
	"github.com/bwjson/fitnessApp/pkg/logger"
	"github.com/bwjson/fitnessApp/pkg/mongodb"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Starting user microservice")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Info("Starting user server")
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Development: %s",
		cfg.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Development,
	)
	appLogger.Infof("Success parsed config: %#v", cfg.AppVersion)

	mongoDBConn, err := mongodb.NewMongoDBConnection(ctx, cfg)
	if err != nil {
		appLogger.Fatal("Cannot connect to MongoDB", err)
	}
	defer func() {
		if err := mongoDBConn.Disconnect(ctx); err != nil {
			appLogger.Fatal("MongoDB disconnection problem: ", err)
		}
	}()
	appLogger.Infof("Connected to MongoDB: %v", mongoDBConn.NumberSessionsInProgress())

	mongoRepo := repository.NewMongoRepository(mongoDBConn)
	services := service.NewService(mongoRepo, appLogger)
	handlers := rest.NewHandler(services, appLogger)

	httpSrv := new(server.HttpServer)

	go func() {
		appLogger.Info("Starting http server", cfg.Http.Port)
		if err := httpSrv.Run(cfg.Http.Port, handlers.InitRoutes()); err != nil {
			appLogger.Fatal("Error starting server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Info("shutting down server...")
}
