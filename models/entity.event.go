package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	ClassID    uint
	Class      Class `gorm:"foreignKey:ClassID"`
	CourseID   uint
	Course     Course `gorm:"foreignKey:CourseID"`
	LecturerID uint
	Lecturer   Lecturer `gorm:"foreignKey:LecturerID"`
	Meet       uint
	Students   []Student `gorm:"many2many:attendances"`
	Status     uint
}
