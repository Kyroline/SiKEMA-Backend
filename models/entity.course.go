package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Code      string
	Name      string
	ClassID   *uint
	Class     Class       `gorm:"foreignKey:ClassID"`
	Lecturers *[]Lecturer `gorm:"many2many:course_lecturers"`
}
