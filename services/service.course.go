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

func (s *CourseService) GetCourseByLecturer() (*[]model.Enrollment, error) {
	//
	var lecturer model.Lecturer
	if err := s.DB.Where("id = ?", 1).Find(&lecturer).Error; err != nil {
		return nil, err
	}
	// s.DB.Find(&lecturer)

	var courses []model.Enrollment
	if err := s.DB.Model(&lecturer).Preload("Class").Preload("Course").Association("Enrollments").Find(&courses); err != nil {
		return nil, err
	}

	return &courses, nil
}

func (s *CourseService) GetCourseByPBM(meta *schema.Metadata) (*[]model.Course, error) {
	var courses []model.Course
	s.DB.Model(&model.Course{}).Preload("Classes").Preload("Lecturers").Count(&meta.Count).Limit(meta.ItemPerPage).Offset((meta.Page - 1) * meta.ItemPerPage).Find(&courses)

	return &courses, nil
}
