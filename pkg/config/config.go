package config

import (
	"injam/delivery/httpserver"
	"injam/repository/postgres"
)

type Config struct {
	Postgres  postgres.Config   `koanf:"postgres"`
	HTTPSever httpserver.Config `koanf:"http_server"`
}
