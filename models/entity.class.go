package model

type Class struct {
	ID       uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name     string    `json:"name,omitempty" gorm:"type:varchar(25);unique"`
	Students []Student `json:"students,omitempty" gorm:"foreignKey:ClassID"`
	Courses  []Course  `json:"courses,omitempty" gorm:"many2many:enrollments"`
}
