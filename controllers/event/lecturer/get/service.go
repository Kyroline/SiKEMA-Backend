package getEvent

import model "attendance-is/models"

type Service interface {
	GetEventService() (*[]model.Event, string)
}

type service struct {
	repository Repository
}

func NewGetEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetEventService() (*[]model.Event, string) {
	event, err := s.repository.GetEvent()
	if err != "" {
		return nil, err
	}
	return event, ""
}
