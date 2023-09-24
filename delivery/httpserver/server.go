package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type Config struct {
	Port uint
}

type Server struct {
	config Config
	Router *echo.Echo
}

func New(config Config) Server {
	return Server{
		config: config,
		Router: echo.New(),
	}
}

func (s Server) Start() {

	s.Router.GET("/health-check", s.healthCheckHandler)

	address := fmt.Sprintf(":%d", s.config.Port)
	if err := s.Router.Start(address); err != nil {
		log.Fatalln("server not start")
	}
}
