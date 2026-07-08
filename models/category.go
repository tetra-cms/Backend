package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Name    string
	Title   string
	IconURL string

	Products []Product `gorm:"foreignKey:CategoryID"`
}
