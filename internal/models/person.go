package models

// Person represent a person Model
type Person struct {
	ID     uint   `json:"ID" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"not null"`
	Gender string `json:"gender" gorm:"not null"`
	Birth  string `json:"birth" gorm:"not null"`
}

func (Person) TableName() string {
	return "persons"
}
