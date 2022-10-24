package middleware

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
