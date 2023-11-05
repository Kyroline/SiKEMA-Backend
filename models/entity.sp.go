package model

type SP struct {
	Id        uint `gorm:"primaryKey"`
	StudentId uint
	Student   Student `gorm:"foreignKey:StudentId"`
	Level     uint
}
