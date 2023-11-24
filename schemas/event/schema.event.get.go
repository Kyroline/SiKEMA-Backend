package schemaEvent

import (
	schemaClass "attendance-is/schemas/class"

	"gorm.io/gorm"
)

type GetEventResponse struct {
	ID         uint
	ClassID    uint                         `json:"-"`
	CourseID   uint                         `json:"-"`
	LecturerID uint                         `json:"-"`
	Class      schemaClass.GetClassResponse `gorm:"foreignKey:ClassID"`
	// Course                    `gorm:"foreignKey:CourseID"`
	// Lecturer   LecturerResponse             `gorm:"foreignKey:LecturerID"`
	Students []StudentResponse `gorm:"many2many:attendances"`
}

type ClassResponse struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type CourseResponse struct {
	ID   uint
	Code string
	Name string
}

type LecturerResponse struct {
	gorm.Model
	Nip  string
	Name string
}

type StudentResponse struct {
	ID   uint `gorm:"primaryKey"`
	Nim  string
	Name string
}
