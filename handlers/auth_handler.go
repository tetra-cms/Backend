package handlers

import (
	"net"
	"net/http"
	"time"

	"tetra-server/auth"
	"tetra-server/config"
	"tetra-server/database"
	"tetra-server/dto"
	"tetra-server/middleware"
	"tetra-server/models"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

// Register godoc
//
//	@Summary		Регистрация пользователя
//	@Description	Создает нового пользователя и возвращает JWT Access Token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RegisterRequest	true	"Данные пользователя"
//	@Success		201		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		409		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var exists models.User

	if err := database.DB.
		Where("username = ?", req.Username).
		First(&exists).Error; err == nil {

		c.JSON(http.StatusConflict, gin.H{
			"error": "username already exists",
		})
		return
	}

	hash, err := auth.HashPassword(req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hash,
		RegIP:    ip,
	}

	if err := database.DB.Create(&user).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := middleware.GenerateAccessToken(&user, h.cfg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

// Login godoc
//
//	@Summary		Авторизация пользователя
//	@Description	Возвращает JWT Access Token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Данные для входа"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	if err := database.DB.
		Where("username = ?", req.Username).
		First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	if err := auth.CheckPassword(user.Password, req.Password); err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := middleware.GenerateAccessToken(&user, h.cfg)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Me godoc
//
//	@Summary		Получить текущего пользователя
//	@Description	Возвращает информацию о текущем авторизованном пользователе
//	@Tags			Auth
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		401	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/auth/me [get]
func (h *AuthHandler) Me(c *gin.Context) {

	userID := c.GetUint("userID")

	var user models.User

	if err := database.DB.
		First(&user, userID).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// Refresh godoc
//
//	@Summary		Обновить Access Token
//	@Description	Возвращает новую пару Access/Refresh токенов
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RefreshRequest	true	"Refresh Token"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Router			/auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {

	var req dto.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	claims, err := middleware.ParseToken(req.RefreshToken, h.cfg)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid refresh token",
		})

		return
	}

	var refresh models.RefreshToken

	err = database.DB.
		Where("token = ?", req.RefreshToken).
		First(&refresh).Error

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "refresh token not found",
		})

		return
	}

	if refresh.Revoked {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "refresh token revoked",
		})

		return
	}

	if refresh.ExpiresAt.Before(time.Now()) {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "refresh token expired",
		})

		return
	}

	var user models.User

	if err := database.DB.First(&user, claims.UserID).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})

		return
	}

	accessToken, _ := middleware.GenerateAccessToken(&user, h.cfg)

	newRefreshToken, _ := middleware.GenerateRefreshToken(&user, h.cfg)

	refresh.Token = newRefreshToken

	refresh.ExpiresAt = time.Now().Add(
		time.Duration(h.cfg.JWTRefreshExpire) * time.Hour,
	)

	database.DB.Save(&refresh)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": newRefreshToken,
	})
}

// Logout godoc
//
//	@Summary		Выход из системы
//	@Description	Отзывает Refresh Token
//	@Tags			Auth
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.RefreshRequest	true	"Refresh Token"
//	@Success		204
//	@Failure		400	{object}	map[string]string
//	@Router			/auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {

	var req dto.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	database.DB.
		Model(&models.RefreshToken{}).
		Where("token = ?", req.RefreshToken).
		Update("revoked", true)

	c.Status(http.StatusNoContent)
}
