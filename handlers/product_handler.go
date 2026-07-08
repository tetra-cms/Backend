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

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// @Summary Получить товары
// @Tags Products
// @Produce json
// @Success 200 {array} dto.ProductResponse
// @Router /products [get]
func (h *ProductHandler) GetAll(c *gin.Context) {

	var products []models.Product

	if err := database.DB.
		Preload("BelongCategory").
		Find(&products).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mapper.ToProductResponses(products))
}

// @Summary Получить товар
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetByID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product

	if err := database.DB.
		Preload("BelongCategory").
		First(&product, id).Error; err != nil {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, mapper.ToProductResponse(&product))
}

// @Summary Создать товар
// @Tags Products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequest true "Product"
// @Success 201 {object} dto.ProductResponse
// @Router /employee/products [post]
func (h *ProductHandler) Create(c *gin.Context) {

	var req dto.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	product := models.Product{
		ImageURL:    req.ImageURL,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	if err := database.DB.Create(&product).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	database.DB.Preload("BelongCategory").First(&product, product.ID)

	c.JSON(http.StatusCreated, mapper.ToProductResponse(&product))
}

// @Summary Обновить товар
// @Tags Products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body dto.UpdateProductRequest true "Product"
// @Success 200 {object} dto.ProductResponse
// @Router /employee/products/{id} [put]
func (h *ProductHandler) Update(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "product not found",
		})
		return
	}

	var req dto.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	product.ImageURL = req.ImageURL
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.CategoryID = req.CategoryID

	if err := database.DB.Save(&product).Error; err != nil {

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	database.DB.Preload("BelongCategory").First(&product, product.ID)

	c.JSON(http.StatusOK, mapper.ToProductResponse(&product))
}

// @Summary Удалить товар
// @Tags Products
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 204
// @Failure 404 {object} dto.ErrorResponse
// @Router /admin/products/{id} [delete]
func (h *ProductHandler) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	result := database.DB.Delete(&models.Product{}, id)

	if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "product not found",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
