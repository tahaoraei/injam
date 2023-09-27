package userserver

import (
	"github.com/labstack/echo/v4"
	"injam/param"
	"log/slog"
	"net/http"
)

func (h Handler) register(c echo.Context) error {
	var req param.UserRegisterRequest
	if err := c.Bind(&req); err != nil {
		slog.Error("error in Binding UserRegisterRequest: ", err)
		return err
	}

	// add claim and jwk
	resp, err := h.userSvc.RegisterUser(c.Request().Context(), req)
	if err != nil {
		slog.Error("error in Register User Service: ", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "can't register"})
	}
	return c.JSON(http.StatusOK, resp)
}
