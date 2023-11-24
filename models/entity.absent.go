package model

import (
	"time"

	"gorm.io/gorm"
)

type Absent struct {
	ID        uint    `gorm:"primaryKey"`
	StudentID uint    `json:"student_id,omitempty"`
	Student   Student `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	EventID   uint
	Event     Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
	Hours     uint
	CreatedAt time.Time
}

func (entity *Absent) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}
