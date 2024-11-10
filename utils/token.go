package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetToken(email string, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userId,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})
	secret := os.Getenv("JWT_SECRET_KEY")
	return token.SignedString([]byte(secret))
}

func VerifyToken() {

}
