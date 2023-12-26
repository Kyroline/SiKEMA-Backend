package model

import (
	"time"

	"gorm.io/gorm"
)

type Lecturer struct {
	UserID uint `json:"user_id,omitempty" gorm:"default:null"`
	gorm.Model
	Nip         string       `json:"nip,omitempty" gorm:"type:varchar(25);unique"`
	Name        string       `json:"name,omitempty" gorm:"type:varchar(255)"`
	Address     *string      `json:"address,omitempty" gorm:"type:varchar(255)"`
	Phone       *string      `json:"phone,omitempty" gorm:"type:varchar(25)"`
	Dob         *time.Time   `json:"dob,omitempty"`
	Enrollments []Enrollment `gorm:"many2many:enrollment_lecturers" json:"enrollments,omitempty"`
}
