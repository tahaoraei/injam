package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"injam/entity"
	"strings"
	"time"
)

const (
	SignKey = "taha"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int
	Name   string
}

type Service struct {
}

func New() Service {
	return Service{}
}

func (s Service) CreateToken(user entity.User, subject string, expireTime time.Duration) (string, error) {
	c := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenStr, err := token.SignedString([]byte(SignKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (s Service) ParseToken(token string) (*Claims, error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, err
}
