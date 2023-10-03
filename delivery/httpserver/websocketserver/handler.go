package websocketserver

import (
	"injam/adapter/redis"
	"injam/service/authorizationservice"
	"injam/service/authservice"
	"injam/service/locationservice"
)

type Handler struct {
	redisAdapter      redis.Adapter
	locationSvc       locationservice.Service
	authenticationSvc authservice.Service
	authorizationSvc  authorizationservice.Service
}

func New(redisAdapter redis.Adapter, locationSvc locationservice.Service, authenticationSvc authservice.Service, authorizationSvc authorizationservice.Service) Handler {
	return Handler{redisAdapter: redisAdapter, locationSvc: locationSvc, authenticationSvc: authenticationSvc, authorizationSvc: authorizationSvc}
}
