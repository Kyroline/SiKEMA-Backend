package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Id      uint
	Name    string    `gorm:"type:varchar(25)"`
	Student []Student `gorm:"foreignKey:ClassId"`
}
