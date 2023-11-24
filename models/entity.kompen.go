package model

import "gorm.io/gorm"

type Kompen struct {
	gorm.Model
	StudentID  uint    `gorm:"foreignKey:StudentID"`
	Student    Student `json:"student,omitempty"`
	Semester   uint
	AmountLeft uint
	Status     bool
}
