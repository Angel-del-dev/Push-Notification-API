package users

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"notificationapi.com/internal/domains/users/dtos"
	"notificationapi.com/internal/infrastructure/request"
)

type Service struct {
	Repository Repository
}

func (s *Service) Store(c fiber.Ctx) error {
	application, found := s.getDataFromRequest(c, "application")
	if !found {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid auth token",
		})
	}

	req, err := request.ParseBody[dtos.RequestStoreType](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	if req.User == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User must be set",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second) // Needed for pooling
	defer cancel()

	exists, err := s.Repository.DoesUserExist(ctx, application, req.User)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	err = s.Repository.CreateUser(ctx, application, req.User)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (s *Service) Remove(c fiber.Ctx) error {
	application, found := s.getDataFromRequest(c, "application")
	if !found {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid auth token",
		})
	}

	req, err := request.ParseBody[dtos.RequestStoreType](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	if req.User == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User must be set",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second) // Needed for pooling
	defer cancel()

	exists, err := s.Repository.DoesUserExist(ctx, application, req.User)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User does not exist",
		})
	}

	err = s.Repository.RemoveUser(ctx, application, req.User)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (s *Service) getDataFromRequest(ctx fiber.Ctx, fieldName string) (string, bool) {
	claims, ok := ctx.Locals("application").(jwt.MapClaims)
	if !ok {
		return "", false
	}

	field, ok := claims[fieldName].(string)
	if !ok {
		return "", false
	}

	return field, true
}
