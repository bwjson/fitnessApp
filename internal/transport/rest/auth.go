package rest

import (
	"github.com/bwjson/fitnessApp/internal/dto"
	"github.com/bwjson/fitnessApp/internal/models"
	"github.com/bwjson/fitnessApp/pkg/http_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register godoc
// @Summary      register
// @Description  creating new account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        registerData body models.User true "Account data"
// @Success      200  {object}  models.User
// @Failure 400 {object} http_errors.HTTPError "Bad Request"
// @Failure 404 {object} http_errors.HTTPError "Not Found"
// @Failure 500 {object} http_errors.HTTPError "Internal Server Error"
// @Router       /auth/register [post]
func (h *Handler) register(c *gin.Context) {
	var data models.User

	ctx := c.Request.Context()

	if err := c.BindJSON(&data); err != nil {
		h.log.Errorf("c.Bind: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	returned_data, err := h.services.Create(ctx, &data)

	// Пробрасываем ошибку из репозитория в кастомные ошибки, а там внутри
	// Выбираем какая это ошибка и создаем HTTP Exception свой
	if err != nil {
		h.log.Errorf("services.Create: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, returned_data)
}

// Login godoc
// @Summary      login
// @Description  authenticate user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginData body dto.LoginInput true "Login data"
// @Success      200  {object}  dto.TokenResponse
// @Failure 400 {object} http_errors.HTTPError "Bad Request"
// @Failure 404 {object} http_errors.HTTPError "Not Found"
// @Failure 500 {object} http_errors.HTTPError "Internal Server Error"
// @Router       /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var data dto.LoginInput
	var tokens dto.TokenResponse

	ctx := c.Request.Context()

	if err := c.BindJSON(&data); err != nil {
		h.log.Errorf("c.Bind: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	tokens, err := h.services.Login(ctx, data)
	if err != nil {
		h.log.Errorf("services.Login: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

// Profile godoc
// @Security ApiKeyAuth
// @Summary      profile info
// @Description  get profile info
// @Tags         user
// @Produce      json
// @Success      200  {object}  models.User
// @Failure 400 {object} http_errors.HTTPError "Bad Request"
// @Failure 404 {object} http_errors.HTTPError "Not Found"
// @Failure 500 {object} http_errors.HTTPError "Internal Server Error"
// @Router       /user/ [get]
func (h *Handler) getProfileInfo(c *gin.Context) {
	ctx := c.Request.Context()

	email, err := getUserEmail(c)
	if err != nil {
		h.log.Errorf("getStudentId: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	user, err := h.services.GetProfileInfo(ctx, email)
	if err != nil {
		h.log.Errorf("services.GetProfileInfo: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"info": user,
	})
}

// ProfileAll godoc
// @Security ApiKeyAuth
// @Summary      all profile info
// @Description  get all profile info
// @Tags         user
// @Produce      json
// @Success      200  {object}  []models.User
// @Failure 400 {object} http_errors.HTTPError "Bad Request"
// @Failure 404 {object} http_errors.HTTPError "Not Found"
// @Failure 500 {object} http_errors.HTTPError "Internal Server Error"
// @Router       /user/all [get]
func (h *Handler) getAll(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.services.GetAllUsers(ctx)
	if err != nil {
		h.log.Errorf("GetAllUsers: %v", err)
		http_errors.NewErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
