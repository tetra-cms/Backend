package mapper

import (
	"tetra-server/dto"
	"tetra-server/models"
)

func ToCategoryResponse(category *models.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:      category.ID,
		Name:    category.Name,
		Title:   category.Title,
		IconURL: category.IconURL,
	}
}

func ToCategoryResponses(categories []models.Category) []dto.CategoryResponse {
	result := make([]dto.CategoryResponse, len(categories))

	for i := range categories {
		result[i] = ToCategoryResponse(&categories[i])
	}

	return result
}
