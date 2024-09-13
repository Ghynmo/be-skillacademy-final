package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	RoleId uint `gorm:"default:1" json:"role_id"`
}



