package model

type Course struct {
	ID        uint   `gorm:"primaryKey"`
	Code      string `gorm:"unique"`
	Name      string
	Classes   *[]Class    `gorm:"many2many:enrollments" json:"classes,omitempty"`
	Lecturers *[]Lecturer `gorm:"many2many:enrollment_lecturers" json:"lecturers,omitempty"`
}
