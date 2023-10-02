package websocketserver

import "github.com/labstack/echo/v4"

func (h Handler) SetRoute(e *echo.Echo) {
	g := e.Group("/ws")
	g.GET("/get_location", h.getLocation)
	g.GET("/share_location:id", h.shareLocation)
}
