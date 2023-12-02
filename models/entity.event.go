package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID         uint           `json:"id,omitempty" gorm:"primaryKey"`
	ClassID    uint           `json:"class_id,omitempty"`
	Class      Class          `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	CourseID   uint           `json:"course_id,omitempty"`
	Course     Course         `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	LecturerID uint           `json:"lecturer_id,omitempty"`
	Lecturer   Lecturer       `json:"lecturer,omitempty" gorm:"foreignKey:LecturerID"`
	Meet       uint           `json:"meet,omitempty"`
	Students   []Student      `json:"students,omitempty" gorm:"many2many:attendances"`
	Status     uint           `json:"status,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
