package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

const (
	GRPC_PORT = ":50051"
	HTTP_PORT = ":8000"
)

type Config struct {
	AppVersion string
	Server     Server
	Http       Http
	MongoDB    MongoDB
	Kafka      Kafka
	Logger     Logger
	AuthConfig AuthConfig
}

type Server struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
	Kafka             Kafka
}

type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type Http struct {
	Port              string
	PprofPort         string
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CookieLifeTime    int
	SessionCookieName string
}

type MongoDB struct {
	URI string
	DB  string
}

type Kafka struct {
	Brokers []string
}

type AuthConfig struct {
	JWT  JWTConfig
	Salt string
}

type JWTConfig struct {
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	SigningKey      string
}

func exportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if os.Getenv("MODE") == "DOCKER" {
		viper.SetConfigName("docker-config.yml")
	} else {
		viper.SetConfigName("config.yml")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func ParseConfig() (*Config, error) {
	if err := exportConfig(); err != nil {
		return nil, err
	}

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode config into string, %v", err)
		return nil, err
	}

	log.Printf("%v", c)

	if c.Server.Port == "" {
		c.Server.Port = GRPC_PORT
	}

	if c.Http.Port == "" {
		c.Http.Port = HTTP_PORT
	}

	return &c, nil
}
