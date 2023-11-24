package model

import (
	"time"

	"gorm.io/gorm"
)

type Lecturer struct {
	UserID *uint `json:"user_id,omitempty"`
	User   User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	gorm.Model
	Nip         string `gorm:"type:varchar(25);unique"`
	Name        string `gorm:"type:varchar(255)"`
	Address     string `gorm:"type:varchar(255)"`
	Dob         time.Time
	Enrollments []Enrollment `gorm:"many2many:enrollment_lecturers" json:"enrollments,omitempty"`
}
