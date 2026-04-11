package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func ContentTypeAllowed(contentTypeAllowed string) func(fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		contentType := ctx.Get("Content-Type")

		if strings.HasPrefix(contentType, contentTypeAllowed) {
			return ctx.Next()
		}
		return ctx.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error": "Content type must be " + contentTypeAllowed,
		})
	}
}
