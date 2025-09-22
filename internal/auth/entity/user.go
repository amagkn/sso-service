package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	Email    string
	PassHash []byte
}

func (u *User) ComparePassword(passHash []byte) error {
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
