package userserver

import (
	"github.com/labstack/echo/v4"
	"injam/param"
	"net/http"
)

func (h Handler) login(c echo.Context) error {
	var req param.UserLoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.validatorSvc.Login(req); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "can't validate user"})
	}

	resp, err := h.userSvc.Login(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cant login user service")
	}

	return c.JSON(http.StatusOK, resp)
}
