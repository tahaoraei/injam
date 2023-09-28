package main

import (
	"injam/delivery/httpserver"
	"injam/pkg/config"
	"injam/pkg/validator"
	"injam/repository/postgres"
	"injam/service/authservice"
	"injam/service/userservice"
)

func main() {
	cfg := config.Load("config.yml")

	repo := postgres.New(cfg.Postgres)
	authSvc := authservice.New()

	userSvc := userservice.New(repo, authSvc)

	validatorSvc := validator.New(repo)

	server := httpserver.New(cfg.HTTPSever, userSvc, validatorSvc, authSvc)
	server.Start()
}
