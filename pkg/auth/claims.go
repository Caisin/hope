package auth

import (
	"context"
	"errors"
	jwt2 "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId   int64  `json:"userId"`
	Avatar   string `json:"avatar"`
	TenantId int64  `json:"tenantId"`
	jwt.StandardClaims
}

func (c Claims) Valid() error {
	err := c.StandardClaims.Valid()
	if err != nil {
		return err
	}
	if c.UserId <= 0 {
		return errors.New("invalid userId")
	}
	if c.TenantId <= 0 {
		return errors.New("invalid tenant")
	}
	return nil
}

func GetClaims(ctx context.Context) (*Claims, error) {
	claims, ok := jwt2.FromContext(ctx)
	if !ok {
		return nil, errors.New("no login")
	}
	c, ok := claims.(Claims)
	if ok {
		return &c, nil
	}
	return nil, errors.New("invalid user type")
}
