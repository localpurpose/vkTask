package models

type Actor struct {
	ID       uint `json:"ID"`
	PersonID uint `json:"person_id" gorm:"index, not null"`
	MovieId  uint `json:"movie_id" gorm:"index, not null"`
}
