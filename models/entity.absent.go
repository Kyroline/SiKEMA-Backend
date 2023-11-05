package model

import "gorm.io/gorm"

type Absent struct {
	gorm.Model
	StudentId uint
	Student   Student `gorm:"foreignKey:StudentId"`
	Hours     uint
}
