package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/httpserver"
)

var _ httpserver.Controller = (*MainController)(nil)

type MainController struct {
	version string
}

func (c *MainController) SetupRoutes(router fiber.Router) {
	router.All("", c.healthHandler)
}

func (c *MainController) healthHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(HealthResponse{Version: c.version})
}

func NewMainController(version string) *MainController {
	return &MainController{version}
}
