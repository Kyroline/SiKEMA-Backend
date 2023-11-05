package model

import "gorm.io/gorm"

type Kompen struct {
	gorm.Model
	StudentId  uint `gorm:"foreignKey:StudentId"`
	Student    Student
	Semester   uint
	AmountLeft uint
	Status     bool
}
