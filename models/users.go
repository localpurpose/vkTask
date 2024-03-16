package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"ID" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null; default:user"`
}
