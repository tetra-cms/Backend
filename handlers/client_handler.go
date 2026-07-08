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

type ClientHandler struct{}

func NewClientHandler() *ClientHandler {
	return &ClientHandler{}
}

// @Summary Получить клиентов текущего пользователя
// @Tags Clients
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.ClientResponse
// @Router /clients [get]
func (h *ClientHandler) GetAll(c *gin.Context) {

	userID := c.GetUint("userID")

	var clients []models.Client

	if err := database.DB.
		Where("user_id = ?", userID).
		Find(&clients).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, mapper.ToClientResponses(clients))
}

// @Summary Получить клиента
// @Tags Clients
// @Security BearerAuth
// @Produce json
// @Param id path int true "Client ID"
// @Success 200 {object} dto.ClientResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /clients/{id} [get]
func (h *ClientHandler) GetByID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	userID := c.GetUint("userID")

	var client models.Client

	err := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&client).Error

	if err != nil {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "client not found",
		})

		return
	}

	c.JSON(http.StatusOK, mapper.ToClientResponse(&client))
}

// @Summary Создать клиента
// @Tags Clients
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateClientRequest true "Client"
// @Success 201 {object} dto.ClientResponse
// @Router /clients [post]
func (h *ClientHandler) Create(c *gin.Context) {

	var req dto.CreateClientRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	client := models.Client{
		UserID:  c.GetUint("userID"),
		FCS:     req.FCS,
		City:    req.City,
		Address: req.Address,
		Phone:   req.Phone,
	}

	if err := database.DB.Create(&client).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, mapper.ToClientResponse(&client))
}

// @Summary Обновить клиента
// @Tags Clients
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Client ID"
// @Param request body dto.UpdateClientRequest true "Client"
// @Success 200 {object} dto.ClientResponse
// @Router /clients/{id} [put]
func (h *ClientHandler) Update(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	userID := c.GetUint("userID")

	var client models.Client

	if err := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&client).Error; err != nil {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "client not found",
		})

		return
	}

	var req dto.UpdateClientRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	client.FCS = req.FCS
	client.City = req.City
	client.Address = req.Address
	client.Phone = req.Phone

	database.DB.Save(&client)

	c.JSON(http.StatusOK, mapper.ToClientResponse(&client))
}

// @Summary Удалить клиента
// @Tags Clients
// @Security BearerAuth
// @Param id path int true "Client ID"
// @Success 204
// @Failure 404 {object} dto.ErrorResponse
// @Router /clients/{id} [delete]
func (h *ClientHandler) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	userID := c.GetUint("userID")

	result := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.Client{})

	if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "client not found",
		})

		return
	}

	c.Status(http.StatusNoContent)
}
