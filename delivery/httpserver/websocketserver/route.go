package websocketserver

import (
	"github.com/labstack/echo/v4"
	"injam/delivery/httpserver/middleware"
)

func (h Handler) SetRoute(e *echo.Echo) {
	g := e.Group("/ws")
	g.GET("/get_location", h.getLocation, middleware.Auth(h.authenticationSvc))
	g.GET("/share_location:id", h.shareLocation, middleware.Auth(h.authenticationSvc), middleware.IsAllowed(h.authorizationSvc))
}
