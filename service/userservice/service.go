package userservice

import (
	"injam/entity"
	"time"
)

type Repository interface {
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phone string) (entity.User, error)
	GetUserByID(id int) (entity.User, error)
}

type AuthGenerator interface {
	CreateToken(user entity.User, subject string, expireTime time.Duration) (string, error)
}

type Service struct {
	repo Repository
	auth AuthGenerator
}

func New(repo Repository, auth AuthGenerator) Service {
	return Service{repo: repo, auth: auth}
}
