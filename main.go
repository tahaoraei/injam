package main

import (
	"injam/delivery/httpserver"
	"injam/pkg/config"
	"injam/repository/postgres"
	"injam/service/userservice"
)

func main() {
	cfg := config.Load("config.yml")

	repo := postgres.New(cfg.Postgres)
	userSvc := userservice.New(repo)

	server := httpserver.New(cfg.HTTPSever, userSvc)
	server.Start()
}
