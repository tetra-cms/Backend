package mapper

import (
	"tetra-server/dto"
	"tetra-server/models"
)

func ToProductResponse(product *models.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:            product.ID,
		ImageURL:      product.ImageURL,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		Stock:         product.Stock,
		SupplyQuantum: product.SupplyQuantum,
		Category:      ToCategoryResponse(&product.Category),
	}
}

func ToProductResponses(products []models.Product) []dto.ProductResponse {
	result := make([]dto.ProductResponse, len(products))

	for i := range products {
		result[i] = ToProductResponse(&products[i])
	}

	return result
}
