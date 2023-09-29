package config

import (
	"injam/adapter/redis"
	"injam/delivery/httpserver"
	"injam/repository/postgres"
)

type Config struct {
	Postgres  postgres.Config   `koanf:"postgres"`
	HTTPSever httpserver.Config `koanf:"http_server"`
	Redis     redis.Config      `koanf:"redis"`
}
