package main

import (
	"context"
	"github.com/bwjson/fitnessApp/config"
	"github.com/bwjson/fitnessApp/internal/server"
	"github.com/bwjson/fitnessApp/pkg/mongodb"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	mongoDBConn, err := mongodb.NewMongoDBConnection(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mongoDBConn.Disconnect(ctx); err != nil {
			log.Fatal("MongoDB disconnection problem: ", err)
		}
	}()

	srv := new(server.HttpServer)

	go func() {
		srv.Run(cfg.Http.Port)
	}()
}
