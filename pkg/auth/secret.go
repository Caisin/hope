package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	JwtKey = "dev"
)

func SecretKeyFun(token *jwt.Token) (interface{}, error) {
	return []byte(JwtKey), nil
}

// ValidatePassword 密码比较 encodePwd:加密的密码 pwd:明文密码
func ValidatePassword(encodePwd string, pwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(pwd))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Encrypt 加密
func Encrypt(pwd string) (string, error) {
	if pwd == "" {
		return pwd, errors.New("password is empty")
	}
	if hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}
