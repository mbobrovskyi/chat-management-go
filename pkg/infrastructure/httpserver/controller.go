package httpserver

import "github.com/gofiber/fiber/v2"

type Controller interface {
	SetupRoutes(router fiber.Router)
}
