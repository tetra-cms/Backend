package dto

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
