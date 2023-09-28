package userserver

import (
	"github.com/labstack/echo/v4"
	"injam/param"
	"injam/service/authservice"
	"net/http"
)

func (h Handler) Profile(c echo.Context) error {
	claims := c.Get("my_claims").(*authservice.Claims)

	resp, err := h.userSvc.Profile(param.UserProfileRequest{ID: claims.UserID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cant get profile")
	}
	return c.JSON(http.StatusOK, resp)
}
