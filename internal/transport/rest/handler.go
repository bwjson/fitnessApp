package rest

import (
	"github.com/bwjson/fitnessApp/internal/service"
	"github.com/bwjson/fitnessApp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	log      logger.Logger
}

func NewHandler(services *service.Service, logger logger.Logger) *Handler {
	return &Handler{services: services, log: logger}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
	}

	return router
}
