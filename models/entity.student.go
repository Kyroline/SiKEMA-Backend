package model

import "gorm.io/gorm"

type Student struct {
	UserId uint
	User   User `gorm:"foreignKey:UserId"`
	gorm.Model
	ClassId uint
	Nim     string
	Name    string
	Events  []Event `gorm:"many2many:student_events"`
}
