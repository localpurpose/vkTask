package models

// Movie represents the model of movie. Rels with actor and persons.
type Movie struct {
	ID          uint   `json:"ID" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null; size:150; unique"`
	Description string `json:"description" gorm:"not null; size:1000"`
	Date        string `json:"date" gorm:"not null"`
	Rating      int    `json:"rating" gorm:"not null"`
}
