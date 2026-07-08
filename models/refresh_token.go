package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model

	Token string `gorm:"uniqueIndex;not null"`

	UserID uint
	User   User

	ExpiresAt time.Time

	Revoked bool `gorm:"default:false"`
}
