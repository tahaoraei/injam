package config

import (
	"injam/delivery/httpserver"
	"injam/repository/postgres"
	"injam/service/userservice"
)

type Config struct {
	Postgres    postgres.Config
	UserService userservice.Config
	HTTPSever   httpserver.Config
}
