package auth

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserId int64  `json:"userId"`
	Avatar string `json:"avatar"`
	jwt.StandardClaims
}
