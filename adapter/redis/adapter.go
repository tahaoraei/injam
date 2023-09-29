package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address  string `koanf:"address"`
	Port     string `koanf:"port"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}

type Adapter struct {
	client *redis.Client
}

func New(config Config) Adapter {
	return Adapter{client: redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Address, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})}
}

func (a Adapter) Client() *redis.Client {
	return a.client
}
