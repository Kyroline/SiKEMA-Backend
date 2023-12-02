package createExcuse

import (
	model "attendance-is/models"
)

type Service interface {
	CreateExcuseService(input *InputCreateExcuse) (*model.Excuse, string)
}

type service struct {
	repository Repository
}

func NewCreateExcuseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateExcuseService(input *InputCreateExcuse) (*model.Excuse, string) {
	newExcuse := model.Excuse{
		AbsentID:   input.ID,
		Excuse:     input.Excuse,
		Attachment: input.Attachment,
		Status:     0,
	}

	result, err := s.repository.CreateExcuseRepository(&newExcuse)

	return result, err
}
