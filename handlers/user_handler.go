package handlers

import (
	"net/http"
	"strconv"

	"tetra-server/database"
	"tetra-server/dto"
	mapper "tetra-server/mappers"
	"tetra-server/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetAll godoc
//
//	@Summary		Получить список пользователей
//	@Description	Возвращает список всех пользователей
//	@Tags			Users
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}		models.User
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/admin/users [get]
func (h *UserHandler) GetAll(c *gin.Context) {

	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, mapper.ToUserResponses(users))
}

// GetByID godoc
//
//	@Summary		Получить пользователя
//	@Tags			Users
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Success		200	{object}	models.User
//	@Failure		404	{object}	dto.ErrorResponse
//	@Router			/admin/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid id",
		})

		return
	}

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "user not found",
		})

		return
	}

	c.JSON(http.StatusOK, mapper.ToUserResponse(&user))
}

// Profile godoc
//
//	@Summary		Профиль пользователя
//	@Tags			Users
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		404	{object}	dto.ErrorResponse
//	@Router			/users/profile [get]
func (h *UserHandler) Profile(c *gin.Context) {

	userID := c.GetUint("userID")

	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "user not found",
		})

		return
	}

	c.JSON(http.StatusOK, mapper.ToUserResponse(&user))
}

// Update godoc
//
//	@Summary		Обновить пользователя
//	@Tags			Users
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int						true	"User ID"
//	@Param			request	body	dto.UpdateUserRequest	true	"User"
//	@Success		200		{object}	models.User
//	@Failure		404		{object}	dto.ErrorResponse
//	@Router			/admin/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})

		return
	}

	user.Username = req.Username
	user.Email = req.Email
	user.Role = req.Role

	database.DB.Save(&user)

	c.JSON(http.StatusOK, mapper.ToUserResponse(&user))
}

// Delete godoc
//
//	@Summary		Удалить пользователя
//	@Tags			Users
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Success		204
//	@Failure		404	{object}	dto.ErrorResponse
//	@Router			/admin/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})

		return
	}

	database.DB.Delete(&user)

	c.Status(http.StatusNoContent)
}
