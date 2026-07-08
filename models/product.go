package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ImageURL string `gorm:"default:'products/no_image.png'"`

	Name        string
	Description string

	Price int

	Stock int `gorm:"default:-1;not null"`

	SupplyQuantum uint `gorm:"default:1;not null"`

	CategoryID uint
	Category   Category
}
