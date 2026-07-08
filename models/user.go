package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`

	Role Role `gorm:"type:varchar(20);default:'USER'"`

	RegIP string

	Clients []Client `gorm:"foreignKey:UserID"`
}
