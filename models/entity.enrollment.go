package model

type Enrollment struct {
	CourseID  uint       `gorm:"primaryKey"`
	ClassID   uint       `gorm:"primaryKey"`
	Course    Course     `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Class     Class      `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Lecturers []Lecturer `gorm:"many2many:enrollment_lecturers"`
}
