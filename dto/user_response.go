package dto

import (
	"tetra-server/models"
	"time"
)

type UserResponse struct {
	ID        uint        `json:"id" example:"1"`
	Username  string      `json:"username" example:"alex"`
	Email     string      `json:"email" example:"alex@example.com"`
	Role      models.Role `json:"role" example:"USER"`
	CreatedAt time.Time   `json:"createdAt"`
}
