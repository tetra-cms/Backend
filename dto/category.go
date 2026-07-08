package dto

type CreateCategoryRequest struct {
	Name    string `json:"name" binding:"required" example:"electronics"`
	Title   string `json:"title" binding:"required" example:"Электроника"`
	IconURL string `json:"iconUrl" binding:"required" example:"categories/electronics.png"`
}

type UpdateCategoryRequest struct {
	Name    string `json:"name" example:"electronics"`
	Title   string `json:"title" example:"Электроника"`
	IconURL string `json:"iconUrl" example:"categories/electronics.png"`
}

type CategoryResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	IconURL string `json:"iconUrl"`
}
