package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
}