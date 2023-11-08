package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(25)"`
	Students []Student `gorm:"foreignKey:ClassId"`
}
