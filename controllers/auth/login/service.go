package loginAuth

import (
	model "attendance-is/models"
)

type Service interface {
	LoginAuthService(input InputLoginAuth) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewLoginAuthService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginAuthService(input InputLoginAuth) (*model.User, string) {
	if input.Email != "" {
		return s.repository.LoginEmailAuthRepository(input.Email, input.Password)
	} else if input.Nim != "" {
		return s.repository.LoginNimAuthRepository(input.Nim, input.Password)
	} else if input.Nip != "" {
		return s.repository.LoginNipAuthRepository(input.Nip, input.Password)
	}
	return nil, "ERR_UNEXPECTED_500"
}
