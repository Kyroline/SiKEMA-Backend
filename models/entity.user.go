package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
	Type     uint
}
