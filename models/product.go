package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ImageURL string `gorm:"default:'products/no_image.png'"`

	Name        string
	Description string

	Price int

	CategoryID uint

	Category Category
}
