package main

import (
	"injam/adapter/redis"
	"injam/delivery/httpserver"
	"injam/pkg/config"
	"injam/pkg/validator"
	"injam/repository/postgres"
	"injam/repository/redis/redislocation"
	"injam/service/authorizationservice"
	"injam/service/authservice"
	"injam/service/locationservice"
	"injam/service/userservice"
)

func main() {
	cfg := config.Load("config.yml")

	repo := postgres.New(cfg.Postgres)
	authSvc := authservice.New()

	userSvc := userservice.New(repo, authSvc)

	validatorSvc := validator.New(repo)

	redisAdapter := redis.New(cfg.Redis)
	inMemoryRepo := redislocation.New(redisAdapter)
	locationSvc := locationservice.New(inMemoryRepo)

	authorizationSvc := authorizationservice.New(repo)

	server := httpserver.New(cfg.HTTPSever, userSvc, validatorSvc, authSvc, authorizationSvc, locationSvc, redisAdapter)
	server.Start()
}
