package rest

import (
	"errors"
	"github.com/bwjson/fitnessApp/pkg/http_errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "email"
)

func (h *Handler) userIdentity(c *gin.Context) {
	email, err := h.parseAuthHeader(c)
	if err != nil {
		http_errors.NewErrorResponse(c, err)
		h.log.Error(email)
		return
	}

	c.Set(userCtx, email)
}

// TODO: Сделать кастомное возвращение ошибок в http_errors
func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("empty auth header")
	}

	return h.tokenManager.Parse(headerParts[1])
}

func getUserEmail(c *gin.Context) (string, error) {
	return getEmailByContext(c, userCtx)
}

func getEmailByContext(c *gin.Context, context string) (string, error) {
	emailFromCtx, ok := c.Get(context)
	if !ok {
		return "", errors.New("userCtx not found")
	}

	email, ok := emailFromCtx.(string)
	if !ok {
		return "", errors.New("userCtx is of invalid type")
	}

	return email, nil
}
