package model

import (
	"time"

	"gorm.io/gorm"
)

type Absent struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	StudentID uint      `json:"student_id,omitempty" gorm:"primaryKey"`
	Student   Student   `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	EventID   uint      `json:"event_id,omitempty" gorm:"primaryKey"`
	Event     Event     `json:"event,omitempty" gorm:"foreignKey:EventID"`
	Hours     uint      `json:"hours,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type AbsentJSON struct {
	ID        uint      `json:"id"`
	StudentID uint      `json:"student_id"`
	Student   *Student  `json:"student,omitempty"`
	EventID   uint      `json:"event_id"`
	Event     *Event    `json:"event,omitempty"`
	Hours     uint      `json:"hours,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (entity *Absent) ParseJSON() *AbsentJSON {
	return &AbsentJSON{
		ID:        entity.ID,
		StudentID: entity.StudentID,
		Student:   &entity.Student,
		EventID:   entity.EventID,
		Event:     &entity.Event,
		Hours:     entity.Hours,
		CreatedAt: entity.CreatedAt,
	}
}

func (entity *Absent) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}
