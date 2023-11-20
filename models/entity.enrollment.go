package model

type Enrollment struct {
	CourseID  uint       `gorm:"primaryKey"`
	ClassID   uint       `gorm:"primaryKey"`
	Lecturers []Lecturer `gorm:"many2many:enrollment_lecturers"`
	Day       string
	Time      string
}
