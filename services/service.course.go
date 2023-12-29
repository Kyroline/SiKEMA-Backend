package service

import (
	model "attendance-is/models"
	schema "attendance-is/schemas"

	"gorm.io/gorm"
)

type InputGetCourseByLecturer struct {
	LecturerID string
}

type CourseService struct {
	DB *gorm.DB
}

func NewCourseService(db *gorm.DB) CourseService {
	return CourseService{DB: db}
}

func (s *CourseService) GetCourseByLecturer(input InputGetCourseByLecturer) (*[]model.Enrollment, error) {
	//
	var lecturer model.Lecturer
	if err := s.DB.Where("id = ?", input.LecturerID).Find(&lecturer).Error; err != nil {
		return nil, err
	}
	// s.DB.Find(&lecturer)

	var courses []model.Enrollment
	if err := s.DB.Model(&lecturer).Preload("Class").Preload("Course").Association("Enrollments").Find(&courses); err != nil {
		return nil, err
	}

	return &courses, nil
}

func (s *CourseService) GetCourseByPBM(meta *schema.Metadata) (*[]model.Enrollment, error) {
	var courses []model.Enrollment
	s.DB.Model(&courses).Preload("Class").Preload("Course").Preload("Lecturers").Find(&courses)

	return &courses, nil
}
