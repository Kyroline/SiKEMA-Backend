package model

import "gorm.io/gorm"

type Excuse struct {
	gorm.Model
	AbsentID   uint
	Absent     Absent `gorm:"foreignKey:AbsentID"`
	Excuse     string `gorm:"type:varchar(255)"`
	Attachment string `gorm:"type:varchar(255)"`
	Status     uint
}
