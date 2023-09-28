package userservice

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"injam/entity"
	"injam/param"
	"log/slog"
)

func (s Service) RegisterUser(ctx context.Context, req param.UserRegisterRequest) (param.UserRegisterResponse, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("cant hash password: ", err)
		return param.UserRegisterResponse{}, err
	}

	u := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    string(pass),
	}

	user, err := s.repo.Register(u)
	if err != nil {
		slog.Error("Error in UserService.Register: ", err)
		return param.UserRegisterResponse{}, err
	}
	return param.UserRegisterResponse{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
	}, nil
}
