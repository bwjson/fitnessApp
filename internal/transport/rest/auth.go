package rest

import (
	"github.com/bwjson/fitnessApp/internal/models"
	"github.com/bwjson/fitnessApp/pkg/http_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) register(c *gin.Context) {
	var data models.User

	ctx := c.Request.Context()

	if err := c.BindJSON(&data); err != nil {
		h.log.Errorf("c.Bind: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	returned_data, err := h.services.Create(ctx, &data)

	if err != nil {
		h.log.Errorf("services.Create: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, returned_data)
}

func (h *Handler) login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
