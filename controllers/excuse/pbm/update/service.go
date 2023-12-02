package updateExcuse

import model "attendance-is/models"

type Service interface {
	UpdateExcuseService(input *InputUpdateExcuse) (*model.Excuse, string)
}

type service struct {
	repository Repository
}

func NewUpdateExcuseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateExcuseService(input *InputUpdateExcuse) (*model.Excuse, string) {
	return nil, ""
}
