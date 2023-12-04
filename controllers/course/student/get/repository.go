package getCourse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetCourseRepository(StudentID string, ClassID string, CourseID string) (*model.Course, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetCourseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetCourseRepository(StudentID string, ClassID string, CourseID string) (*model.Course, string) {
	var class model.Class
	r.db.Where("id = ?", ClassID).Find(&class)

	var courses model.Course
	if err := r.db.Model(&class).Where("id = ?", CourseID).Association("Courses").Find(&courses); err != nil {
		return nil, "500_INTERNAL"
	}

	if courses.ID == 0 {
		return nil, "404_NOTFOUND"
	}

	return &courses, ""
}
