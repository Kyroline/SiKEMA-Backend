package model

import "gorm.io/gorm"

type Absent struct {
	gorm.Model
	StudentID uint
	Student   Student `gorm:"foreignKey:StudentID"`
	EventID   uint
	Event     Event `gorm:"foreignKey:EventID"`
	Hours     uint
}
