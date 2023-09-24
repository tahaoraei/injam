package main

import "injam/delivery/httpserver"

func main() {
	cfg := httpserver.Config{Port: 8088}
	server := httpserver.New(cfg)
	server.Start()
}
