package getCourse

import (
	model "attendance-is/models"
)

type Service interface {
	GetCourseService(input InputGetCourse) (*model.Course, string)
}

type service struct {
	repository Repository
}

func NewGetCourseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCourseService(input InputGetCourse) (*model.Course, string) {
	res, err := s.repository.GetCourseRepository(input.StudentID, input.ClassID, input.CourseID)

	return res, err
}
