package userservice

import (
	"injam/entity"
)

type Repository interface {
	Register(user entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
