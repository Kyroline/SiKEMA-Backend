package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	ClassId    uint
	Class      Class `gorm:"foreignKey:ClassId"`
	LecturerId uint
	Lecturer   Lecturer `gorm:"foreignKey:LecturerId"`
	Meet       uint
	Students   []Student `gorm:"many2many:student_events"`
	Status     uint
}
