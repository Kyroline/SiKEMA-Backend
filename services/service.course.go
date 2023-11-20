package service

import (
	model "attendance-is/models"
	schema "attendance-is/schemas"

	"gorm.io/gorm"
)

type CourseService struct {
	DB *gorm.DB
}

func NewCourseService(db *gorm.DB) CourseService {
	return CourseService{DB: db}
}

func (s *CourseService) GetCourseByLecturer() (*[]model.Course, error) {
	//
	var lecturer model.Lecturer
	// s.DB.Where("id = ?", 1).Find(&lecturer)
	s.DB.Find(&lecturer)

	var courses []model.Course
	s.DB.Model(&lecturer).Association("Courses").Find(&courses)

	return &courses, nil
}

func (s *CourseService) GetCourseByPBM(meta schema.Metadata) (*[]model.Course, error) {
	var courses []model.Course
	s.DB.Model(&model.Course{}).Limit(meta.ItemPerPage).Offset((meta.Page - 1) * meta.ItemPerPage).Find(&courses)

	return &courses, nil
}
