package userserver

import "injam/service/userservice"

type Handler struct {
	userSvc userservice.Service
}

func New(userSvc userservice.Service) Handler {
	return Handler{
		userSvc: userSvc,
	}
}
