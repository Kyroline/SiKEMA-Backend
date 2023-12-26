package loginAuth

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"strconv"
)

type Service interface {
	LoginAuthService(input InputLoginAuth) (string, string)
}

type service struct {
	repository Repository
}

func NewLoginAuthService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginAuthService(input InputLoginAuth) (string, string) {
	var user model.User
	if input.Email != "" {
		res, err := s.repository.LoginEmailAuthRepository(input.Email, input.Password)
		if err != "" {
			return "", err
		}
		user = *res
	} else if input.Nim != "" {
		res, err := s.repository.LoginNimAuthRepository(input.Nim, input.Password)
		if err != "" {
			return "", err
		}
		user = *res
	} else if input.Nip != "" {
		res, err := s.repository.LoginNipAuthRepository(input.Nip, input.Password)
		if err != "" {
			return "", err
		}
		user = *res
	} else {
		return "", "ERR_UNAUTHENTICATED"
	}
	var typeID uint
	if user.Lecturer != nil {
		typeID = user.Lecturer.ID
	}
	if user.Student != nil {
		typeID = user.Student.ID
	}
	claim := util.JWTClaim{
		Uid:    strconv.FormatUint(uint64(user.ID), 10),
		Type:   strconv.FormatUint(uint64(user.Type), 10),
		TypeID: strconv.FormatUint(uint64(typeID), 10),
	}
	token, err := util.GenerateToken(claim)
	if err != nil {
		return "", "TOKEN_ERR_" + err.Error()
	}
	return token, ""
}
