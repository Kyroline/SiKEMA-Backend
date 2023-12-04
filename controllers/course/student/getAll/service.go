package getAllCourse

import (
	model "attendance-is/models"
)

type Service interface {
	GetAllCourseService(input InputGetAllCourse) (*[]model.Course, string)
}

type service struct {
	repository Repository
}

func NewGetAllCourseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAllCourseService(input InputGetAllCourse) (*[]model.Course, string) {
	res, err := s.repository.GetAllCourseRepository(input.StudentID, input.ClassID)

	return res, err
}