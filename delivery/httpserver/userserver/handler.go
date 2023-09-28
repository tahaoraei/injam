package userserver

import (
	"injam/pkg/validator"
	"injam/service/userservice"
)

type Handler struct {
	userSvc      userservice.Service
	validatorSvc validator.Validator
}

func New(userSvc userservice.Service, validatorSvc validator.Validator) Handler {
	return Handler{
		userSvc:      userSvc,
		validatorSvc: validatorSvc,
	}
}
