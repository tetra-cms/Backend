package dto

type CreateClientRequest struct {
	FCS     string `json:"fcs" binding:"required" example:"Иванов Иван Иванович"`
	City    string `json:"city" binding:"required" example:"Москва"`
	Address string `json:"address" binding:"required" example:"ул. Ленина, 1"`
	Phone   string `json:"phone" binding:"required" example:"+79991234567"`
}

type UpdateClientRequest struct {
	FCS     string `json:"fcs" example:"Иванов Иван Иванович"`
	City    string `json:"city" example:"Москва"`
	Address string `json:"address" example:"ул. Ленина, 1"`
	Phone   string `json:"phone" example:"+79991234567"`
}

type ClientResponse struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"userId"`
	FCS     string `json:"fcs"`
	City    string `json:"city"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
