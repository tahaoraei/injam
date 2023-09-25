package userserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) register(c echo.Context) error {
	return c.String(http.StatusOK, "what do you want?")
}
