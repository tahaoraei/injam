package main

import (
	"injam/delivery/httpserver"
	"injam/service/userservice"
)

func main() {
	cfg := httpserver.Config{Port: 8088}

	userSvc := userservice.New()

	server := httpserver.New(cfg, userSvc)
	server.Start()
}
