package auth

import (
	"context"
	"errors"
	jwt2 "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
)

type Claims struct {
	UserId   int64  `json:"userId,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	TenantId int64  `json:"tenantId,omitempty"`
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
	mc, ok := claims.(jwt.MapClaims)
	if ok {
		//todo 临时解决方案
		c = Claims{}
		c.UserId = cast.ToInt64(mc["userId"])
		c.Avatar = cast.ToString(mc["avatar"])
		c.TenantId = cast.ToInt64(mc["tenantId"])
		c.Audience = cast.ToString(mc["aud"])
		c.ExpiresAt = cast.ToInt64(mc["exp"])
		c.Id = cast.ToString(mc["jti"])
		c.IssuedAt = cast.ToInt64(mc["iat"])
		c.Issuer = cast.ToString(mc["iss"])
		c.NotBefore = cast.ToInt64(mc["nbf"])
		c.Subject = cast.ToString(mc["sub"])
		/*err := c.Valid()
		if err != nil {
			return nil, err
		}*/
		return &c, nil
	}
	return nil, errors.New("invalid user type")
}
