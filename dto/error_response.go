package dto

type ErrorResponse struct {
	Error string `json:"error" example:"user not found"`
}
