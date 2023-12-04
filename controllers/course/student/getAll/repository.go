package getAllCourse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllCourseRepository(StudentID string, ClassID string) (*[]model.Enrollment, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetAllCourseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllCourseRepository(StudentID string, ClassID string) (*[]model.Enrollment, string) {
	// var class model.Class
	// r.db.Where("id = ?", ClassID).Find(&class)

	// var courses []model.Course
	// if err := r.db.Model(&class).Association("Courses").Find(&courses); err != nil {
	// 	return nil, "500_INTERNAL"
	// }

	// if len(courses) == 0 {
	// 	return nil, "404_NOTFOUND"
	// }

	var courses []model.Enrollment
	r.db.Model(model.Enrollment{}).Preload("Course").Preload("Lecturers").Where("class_id = ?", ClassID).Find(&courses)

	return &courses, ""
}
