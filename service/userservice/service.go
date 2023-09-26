package userservice

import (
	"context"
	"log/slog"
)

type Config struct {
}

type Repository interface {
	Register(ctx context.Context, phoneNumber string) error
}

type Service struct {
	config Config
	repo   Repository
}

func New(config Config, repo Repository) Service {
	return Service{
		config: config,
		repo:   repo,
	}
}

func (s Service) RegisterUser(ctx context.Context, phoneNumber string) error {
	if err := s.repo.Register(ctx, phoneNumber); err != nil {
		slog.Error("Error in UserService.Register: ", err)
		return err
	}
	return nil
}
