package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID         uint     `gorm:"primaryKey"`
	ClassID    uint     `json:"class_id,omitempty"`
	Class      Class    `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	CourseID   uint     `json:"course_id,omitempty"`
	Course     Course   `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	LecturerID uint     `json:"lecturer_id,omitempty"`
	Lecturer   Lecturer `gorm:"foreignKey:LecturerID" json:"lecturer,omitempty"`
	Meet       uint
	Students   []Student `gorm:"many2many:attendances"`
	Status     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
