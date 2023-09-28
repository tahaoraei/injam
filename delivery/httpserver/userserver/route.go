package userserver

import (
	"github.com/labstack/echo/v4"
	"injam/delivery/httpserver/middleware"
)

func (h Handler) SetRoute(e *echo.Echo) {
	g := e.Group("/users")

	// check validation with middleware
	g.POST("/register", h.register)
	g.POST("/login", h.login)
	g.GET("/profile", h.Profile, middleware.Auth(h.authSvc))
}
