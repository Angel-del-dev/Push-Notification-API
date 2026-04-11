package auth

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"notificationapi.com/internal/domains/auth/dtos"
	"notificationapi.com/internal/infrastructure/request"
)

type Service struct {
	Repository Repository
	Secret     string
}

func (s *Service) Login(c fiber.Ctx) error {
	req, err := request.ParseBody[dtos.LoginType](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	if req.Application == "" || req.Key == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing authentication information",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second) // Needed for pooling
	defer cancel()

	application, err := s.Repository.GetApplicationByDomain(ctx, req.Application, req.Key)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(application.Password),
		[]byte(req.Password),
	); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, exp, err := s.generateToken(application.Application, application.Key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot generate JWT Token",
		})
	}
	expiresIn := int(time.Until(exp).Seconds())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
		"expires_at":   exp,
		"expires_in":   expiresIn,
	})
}

func (s *Service) generateToken(application string, key string) (string, time.Time, error) {
	expiration := time.Now().Add(time.Minute * 15)
	claims := jwt.MapClaims{
		"application": application,
		"key":         key,
		"exp":         expiration.Unix(),
		"iat":         time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.Secret))
	return signed, expiration, err
}
