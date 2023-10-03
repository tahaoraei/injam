package middleware

import (
	"github.com/labstack/echo/v4"
	"injam/service/authorizationservice"
	"injam/service/authservice"
	"net/http"
	"strconv"
	"strings"
)

func IsAllowed(authorizationSvc authorizationservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("my_claims").(*authservice.Claims)
			pub_id, err := strconv.Atoi(strings.Replace(c.Param("id"), ":", "", 1))
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			isAllowed, err := authorizationSvc.CheckPermission(pub_id, claims.UserID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			if !isAllowed {
				return c.JSON(http.StatusUnauthorized, "you don't have access to this user")
			}

			return next(c)
		}
	}
}
