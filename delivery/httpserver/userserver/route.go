package userserver

import "github.com/labstack/echo/v4"

func (h Handler) SetRoute(e *echo.Echo) {
	g := e.Group("/users")
	g.POST("/register", h.register)

}
