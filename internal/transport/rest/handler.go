package rest

import (
	_ "github.com/bwjson/fitnessApp/docs"
	"github.com/bwjson/fitnessApp/internal/service"
	"github.com/bwjson/fitnessApp/pkg/auth"
	"github.com/bwjson/fitnessApp/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services     *service.Service
	log          logger.Logger
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Service, logger logger.Logger, tokenManager *auth.TokenManager) *Handler {
	return &Handler{services: services, log: logger, tokenManager: *tokenManager}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)

	}

	user := router.Group("/user", h.userIdentity)
	{
		// Берем почту из хэдеров через мидлвэйр слой
		user.GET("/", h.getProfileInfo)
		user.GET("/all", h.getAll)
	}

	return router
}
