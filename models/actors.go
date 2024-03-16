package models

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	ID       uint `json:"ID"`
	PersonID int  `json:"person_id" gorm:"index, not null"`
	MovieId  int  `json:"movie_id" gorm:"index, not null"`
}
