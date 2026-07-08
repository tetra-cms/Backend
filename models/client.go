package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	UserID uint

	User User

	FCS     string
	City    string
	Address string
	Phone   string
}
