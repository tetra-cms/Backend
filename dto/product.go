package dto

type CreateProductRequest struct {
	ImageURL    string `json:"imageUrl" example:"products/iphone.png"`
	Name        string `json:"name" binding:"required" example:"iPhone 16 Pro"`
	Description string `json:"description" binding:"required" example:"Новый смартфон Apple"`
	Price       int    `json:"price" binding:"required" example:"129990"`
	CategoryID  uint   `json:"categoryId" binding:"required" example:"1"`
}

type UpdateProductRequest struct {
	ImageURL    string `json:"imageUrl"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	CategoryID  uint   `json:"categoryId"`
}

type ProductResponse struct {
	ID          uint             `json:"id"`
	ImageURL    string           `json:"imageUrl"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       int              `json:"price"`
	Category    CategoryResponse `json:"category"`
}
