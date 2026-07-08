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

type CategoryHandler struct{}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

// @Summary Получить список категорий
// @Tags Categories
// @Produce json
// @Success 200 {array} dto.CategoryResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /categories [get]
func (h *CategoryHandler) GetAll(c *gin.Context) {

	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, mapper.ToCategoryResponses(categories))
}

// @Summary Получить категорию
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} dto.CategoryResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid id",
		})

		return
	}

	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "category not found",
		})

		return
	}

	c.JSON(http.StatusOK, mapper.ToCategoryResponse(&category))
}

// @Summary Создать категорию
// @Tags Categories
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateCategoryRequest true "Category"
// @Success 201 {object} dto.CategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /admin/categories [post]
func (h *CategoryHandler) Create(c *gin.Context) {

	var req dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	category := models.Category{
		Name:    req.Name,
		Title:   req.Title,
		IconURL: req.IconURL,
	}

	if err := database.DB.Create(&category).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, mapper.ToCategoryResponse(&category))
}

// @Summary Обновить категорию
// @Tags Categories
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param request body dto.UpdateCategoryRequest true "Category"
// @Success 200 {object} dto.CategoryResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /admin/categories/{id} [put]
func (h *CategoryHandler) Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid id",
		})

		return
	}

	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "category not found",
		})

		return
	}

	var req dto.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	category.Name = req.Name
	category.Title = req.Title
	category.IconURL = req.IconURL

	if err := database.DB.Save(&category).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, mapper.ToCategoryResponse(&category))
}

// @Summary Удалить категорию
// @Tags Categories
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 204
// @Failure 404 {object} dto.ErrorResponse
// @Router /admin/categories/{id} [delete]
func (h *CategoryHandler) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid id",
		})

		return
	}

	result := database.DB.Delete(&models.Category{}, id)

	if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "category not found",
		})

		return
	}

	c.Status(http.StatusNoContent)
}
