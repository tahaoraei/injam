package userserver

import (
	"injam/pkg/validator"
	"injam/service/authservice"
	"injam/service/userservice"
)

type Handler struct {
	userSvc      userservice.Service
	validatorSvc validator.Validator
	authSvc      authservice.Service
}

func New(userSvc userservice.Service, validatorSvc validator.Validator, authSvc authservice.Service) Handler {
	return Handler{
		userSvc:      userSvc,
		validatorSvc: validatorSvc,
		authSvc:      authSvc,
	}
}
