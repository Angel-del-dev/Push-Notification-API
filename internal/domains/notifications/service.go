package notifications

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"notificationapi.com/internal/infrastructure/request"
	"notificationapi.com/pkg"
)

type NotificationService struct {
	PublicKey  string
	PrivateKey string
}

func (s *NotificationService) GenerateVAPIDKeys(ctx fiber.Ctx) error {
	vapidPublicKey, vapidPrivateKey, err := pkg.GenerateVAPIDKeys()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate VAPID keys",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"vapidPublicKey":  vapidPublicKey,
		"vapidPrivateKey": vapidPrivateKey,
	})
}

func (s *NotificationService) CheckVAPIDKeys(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"vapidPublicKey":  s.PublicKey,
		"vapidPrivateKey": s.PrivateKey,
	})
}

func (s *NotificationService) Subscribe(ctx fiber.Ctx) error {

	req, err := request.ParseBody[PushSubscription](ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// 👇 Aquí guardas en DB (ejemplo simple)
	fmt.Println("Nueva suscripción:")
	fmt.Println("Endpoint:", req.Endpoint)
	fmt.Println("P256dh:", req.Keys.P256dh)
	fmt.Println("Auth:", req.Keys.Auth)

	return ctx.JSON(fiber.Map{
		"message": "Subscription saved",
	})
}
