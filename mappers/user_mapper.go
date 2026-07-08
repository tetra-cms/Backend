package mapper

import (
	"tetra-server/dto"
	"tetra-server/models"
)

func ToUserResponse(user *models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}

func ToUserResponses(users []models.User) []dto.UserResponse {
	response := make([]dto.UserResponse, len(users))

	for i := range users {
		response[i] = ToUserResponse(&users[i])
	}

	return response
}
