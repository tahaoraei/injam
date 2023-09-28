package userservice

import (
	"golang.org/x/crypto/bcrypt"
	"injam/param"
	"time"
)

func (s Service) Login(req param.UserLoginRequest) (param.UserLoginResponse, error) {
	u, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.UserLoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return param.UserLoginResponse{}, err
	}

	token, err := s.auth.CreateToken(u, "access", time.Hour)
	if err != nil {
		return param.UserLoginResponse{}, err
	}

	return param.UserLoginResponse{
		ID:    u.ID,
		Name:  u.Name,
		Token: token,
	}, nil
}
