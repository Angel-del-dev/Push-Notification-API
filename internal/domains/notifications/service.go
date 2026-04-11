package notifications

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"notificationapi.com/internal/infrastructure/request"
	"notificationapi.com/pkg"
)

type Service struct {
	Repository Repository
	PublicKey  string
	PrivateKey string
}

func (s *Service) GenerateVAPIDKeys(ctx fiber.Ctx) error {
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

func (s *Service) CheckVAPIDKeys(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"vapidPublicKey":  s.PublicKey,
		"vapidPrivateKey": s.PrivateKey,
	})
}

func (s *Service) Subscribe(ctx fiber.Ctx) error {

	req, err := request.ParseBody[PushSubscription](ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	fmt.Println("Nueva suscripción:")
	fmt.Println("Endpoint:", req.Endpoint)
	fmt.Println("P256dh:", req.Keys.P256dh)
	fmt.Println("Auth:", req.Keys.Auth)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Subscription saved",
	})
}

func (s *Service) Send(ctx fiber.Ctx) error {
	var subscription pkg.StoredSubscription

	req, err := request.ParseBody[PushSubscription](ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	subscription.Endpoint = req.Endpoint
	subscription.Auth = req.Keys.Auth
	subscription.P256dh = req.Keys.P256dh
	fmt.Println(subscription)

	payload := map[string]string{
		"title": "Hola 👋",
		"body":  "Notificación enviada desde Go 🚀",
	}

	statuscode, err := pkg.SendNotification(subscription, s.PublicKey, s.PrivateKey, payload)

	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	fmt.Println("Statuscode: ")
	fmt.Println(statuscode)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Notification Sent",
	})
}
