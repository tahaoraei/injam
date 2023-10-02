package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"injam/adapter/redis"
	"injam/delivery/httpserver/userserver"
	"injam/delivery/httpserver/websocketserver"
	"injam/pkg/validator"
	"injam/service/authservice"
	"injam/service/locationservice"
	"injam/service/userservice"
	"log"
)

type Config struct {
	Port uint `koanf:"port"`
}

type Server struct {
	config       Config
	userHandler  userserver.Handler
	odHandler    websocketserver.Handler
	validatorSvc validator.Validator
	Router       *echo.Echo
}

func New(
	config Config,
	userSvc userservice.Service,
	validatorSvc validator.Validator,
	authSvc authservice.Service,
	locationSvc locationservice.Service,
	redisAdapter redis.Adapter,
) Server {
	return Server{
		config:       config,
		userHandler:  userserver.New(userSvc, validatorSvc, authSvc),
		odHandler:    websocketserver.New(redisAdapter, locationSvc),
		validatorSvc: validatorSvc,
		Router:       echo.New(),
	}
}

func (s Server) Start() {
	s.Router.GET("/health-check", s.healthCheckHandler)

	s.userHandler.SetRoute(s.Router)
	s.odHandler.SetRoute(s.Router)

	address := fmt.Sprintf(":%d", s.config.Port)
	if err := s.Router.Start(address); err != nil {
		log.Fatalln("server not start")
	}
}
