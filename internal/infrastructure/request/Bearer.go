package request

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(c fiber.Ctx, secret []byte) (jwt.MapClaims, bool) {
	tokenStr := c.Get("Authorization")
	const bearerPrefix = "Bearer "
	if strings.HasPrefix(tokenStr, bearerPrefix) {
		tokenStr = strings.TrimPrefix(tokenStr, bearerPrefix)
	} else {
		return nil, false
	}

	if tokenStr == "" {
		return nil, false
	}

	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}

	return claims, true
}
