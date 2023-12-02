package model

import (
	"time"

	"gorm.io/gorm"
)

type Excuse struct {
	ID         uint      `gorm:"primaryKey" json:"id,omitempty"`
	AbsentID   uint      `json:"absent_id,omitempty"`
	Absent     Absent    `gorm:"foreignKey:AbsentID" json:"absent,omitempty"`
	Excuse     string    `gorm:"type:varchar(255)" json:"excuse,omitempty"`
	Attachment string    `gorm:"type:varchar(255)" json:"attachment,omitempty"`
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
