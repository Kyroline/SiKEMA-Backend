package model

import (
	"time"
)

type Student struct {
	UserID  *uint
	User    User `gorm:"foreignKey:UserID" json:"-"`
	ID      uint `gorm:"primaryKey"`
	ClassID uint
	Class   Class  `gorm:"foreignKey:ClassID" json:"-"`
	Nim     string `gorm:"type:varchar(25)"`
	Name    string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`
	Dob     time.Time
	Year    uint
}
