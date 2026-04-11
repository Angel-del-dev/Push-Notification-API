package router

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"notificationapi.com/internal/config"
	"notificationapi.com/internal/domains/notifications"
	"notificationapi.com/internal/infrastructure/service"
)

type Router struct {
	app           *fiber.App
	configuration *config.Config
}

func (r *Router) Initialize() {
	r.setDefaultVariables()
	r.createRouter()
	err := r.app.Listen(":" + r.configuration.Application.Port)
	if err != nil {
		log.Panicf("Failed to start server: %v", err)
	}
}

func (r *Router) setDefaultVariables() {
	r.configuration = &config.Config{}
	r.configuration.Application.MaxRequestsPerMinute = 1000
	r.configuration.Application.Port = "3000"
	r.configuration.Application.VAPIDPrivateKey = os.Getenv("VAPIDPRIVATEKEY")
	r.configuration.Application.VAPIDPubliKey = os.Getenv("VAPIDPUBLICKEY")

	r.app = fiber.New(fiber.Config{
		Immutable: true,
	})

	r.setRateLimiter()
}

func (r *Router) setRateLimiter() {
	if r.configuration.Application.MaxRequestsPerMinute == 0 {
		return
	}
	r.app.Use(limiter.New(limiter.Config{
		Max:        r.configuration.Application.MaxRequestsPerMinute,
		Expiration: 1 * time.Minute,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	}))
}

func (r *Router) createRouter() {
	r.setWebPush()
}

func (r *Router) setWebPush() {
	notificationService := service.NewService[notifications.NotificationService]()
	notificationService.PrivateKey = r.configuration.Application.VAPIDPrivateKey
	notificationService.PublicKey = r.configuration.Application.VAPIDPubliKey

	webpushGroup := r.app.Group("/notifications")
	/*
		// Config only
		webpushGroup.Get("/vapid-keys", notificationService.GenerateVAPIDKeys)
		webpushGroup.Get("/CheckVAPIDKeys", notificationService.CheckVAPIDKeys)
	*/
	webpushGroup.Post("/subscribe", notificationService.Subscribe)
}
