package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"injam/delivery/httpserver/userserver"
	"injam/service/userservice"
	"log"
)

type Config struct {
	Port uint `koanf:"port"`
}

type Server struct {
	config      Config
	userHandler userserver.Handler
	Router      *echo.Echo
}

func New(config Config, userSvc userservice.Service) Server {
	return Server{
		config:      config,
		userHandler: userserver.New(userSvc),
		Router:      echo.New(),
	}
}

func (s Server) Start() {
	s.Router.GET("/health-check", s.healthCheckHandler)
	s.userHandler.SetRoute(s.Router)

	address := fmt.Sprintf(":%d", s.config.Port)
	if err := s.Router.Start(address); err != nil {
		log.Fatalln("server not start")
	}
}
