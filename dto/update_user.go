package dto

import "tetra-server/models"

type UpdateUserRequest struct {
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Role     models.Role `json:"role"`
}
