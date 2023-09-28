package middleware

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"injam/service/authservice"
)

func Auth(authSvc authservice.Service) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey:    "my_claims",
		SigningKey:    "taha",
		SigningMethod: "HS256",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			return authSvc.ParseToken(auth)
		},
	})
}
