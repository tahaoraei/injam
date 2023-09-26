package userserver

import "github.com/labstack/echo/v4"

func (h Handler) SetRoute(e *echo.Echo) {
	g := e.Group("/users")

	// check validation with middleware
	g.POST("/register", h.register)
	g.POST("/login", h.login)
}
