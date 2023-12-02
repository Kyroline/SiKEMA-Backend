package model

import (
	"time"

	"gorm.io/gorm"
)

type Absent struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	StudentID uint      `json:"student_id,omitempty"`
	Student   Student   `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	EventID   uint      `json:"event_id,omitempty"`
	Event     Event     `json:"event,omitempty" gorm:"foreignKey:EventID"`
	Hours     uint      `json:"hours,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (entity *Absent) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}
