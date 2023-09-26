package userserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) login(c echo.Context) error {
	return c.String(http.StatusOK, "login :)")
}
