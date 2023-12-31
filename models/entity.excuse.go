package model

import (
	"time"

	"gorm.io/gorm"
)

type Excuse struct {
	ID         uint      `json:"id,omitempty" gorm:"primaryKey"`
	AbsentID   uint      `json:"absent_id,omitempty"`
	Absent     *Absent   `json:"absent,omitempty" gorm:"foreignKey:AbsentID"`
	Excuse     string    `json:"excuse,omitempty" gorm:"type:varchar(255)"`
	Attachment string    `json:"attachment,omitempty" gorm:"type:varchar(255)"`
	Status     uint      `json:"status,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

func (entity *Excuse) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Excuse) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
