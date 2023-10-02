package websocketserver

import (
	"injam/adapter/redis"
	"injam/service/locationservice"
)

type Handler struct {
	redisAdapter redis.Adapter
	locationSvc  locationservice.Service
}

func New(redisAdapter redis.Adapter, locationSvc locationservice.Service) Handler {
	return Handler{redisAdapter: redisAdapter, locationSvc: locationSvc}
}
