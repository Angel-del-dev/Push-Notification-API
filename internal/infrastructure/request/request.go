package request

import "github.com/gofiber/fiber/v3"

func ParseBody[T any](c fiber.Ctx) (T, error) {
	var body T
	err := c.Bind().Body(&body)
	if err != nil {
		return body, err
	}

	return body, nil
}
