package model

import (
	"time"
)

type Student struct {
	UserID  *uint     `json:"user_id,omitempty"`
	User    User      `gorm:"foreignKey:UserID" json:"-"`
	ID      uint      `gorm:"primaryKey"`
	ClassID uint      `json:"class_id,omitempty"`
	Class   Class     `gorm:"foreignKey:ClassID" json:"-"`
	Nim     string    `gorm:"type:varchar(25);unique"`
	Name    string    `gorm:"type:varchar(255)"`
	Address string    `gorm:"type:varchar(255)"`
	Dob     time.Time `gorm:"type:date"`
	Year    uint
}
