package entity

import (
	"fmt"
	"time"

	"github.com/amagkn/sso-service/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	Email    string
	PassHash []byte
	IsAdmin  bool
}

func (u *User) ComparePassword(passHash []byte) error {
	logger.Info(fmt.Sprintf("%s, %s", u.PassHash, passHash))
	return bcrypt.CompareHashAndPassword(u.PassHash, passHash)
}

func (u *User) NewJWTToken(app App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = u.ID
	claims["email"] = u.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
