package models

import (
	"gorm.io/gorm"
)

// Movie represents the model
type Movie struct {
	// gorm.Model
	gorm.Model
	ID          uint   `json:"ID" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null; size:150; unique"`
	Description string `json:"description" gorm:"not null; size:1000"`
	Date        string `json:"date" gorm:"not null"`
	Rating      int    `json:"rating" gorm:"not null"`
}
