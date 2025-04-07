package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "missing authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"error":   true,
				"message": "invalid token format",
			})
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"error":   true,
				"message": "invalid or expired token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user", claims)
		}

		return next(c)
	}
}
