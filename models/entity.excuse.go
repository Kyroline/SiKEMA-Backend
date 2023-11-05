package model

import "gorm.io/gorm"

type Excuse struct {
	gorm.Model
	AbsentId   uint
	Absent     Absent `gorm:"foreignKey:AbsentId"`
	Excuse     string `gorm:"type:varchar(255)"`
	Attachment string `gorm:"type:varchar(255)"`
	Status     bool
}
