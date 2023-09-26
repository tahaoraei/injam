package main

import (
	"injam/delivery/httpserver"
	"injam/pkg/config"
	"injam/repository/postgres"
	"injam/service/userservice"
)

func main() {
	cfg := config.Config{
		Postgres:    postgres.Config{},
		UserService: userservice.Config{},
		HTTPSever:   httpserver.Config{Port: 8088},
	}

	repo := postgres.New(cfg.Postgres)
	userSvc := userservice.New(cfg.UserService, repo)

	server := httpserver.New(cfg.HTTPSever, userSvc)
	server.Start()
}
