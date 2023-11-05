package model

import "gorm.io/gorm"

type Lecturer struct {
	UserId uint
	User   User `gorm:"foreignKey:UserId"`
	gorm.Model
	Nip  string
	Name string
}
