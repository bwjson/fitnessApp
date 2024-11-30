package rest

import "github.com/gin-gonic/gin"

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {

		return
	}
}
