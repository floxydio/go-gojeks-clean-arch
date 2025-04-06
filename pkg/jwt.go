package pkg

import (
	_ "fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"

	"time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(name string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":    name,
			"user_id": userId,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
