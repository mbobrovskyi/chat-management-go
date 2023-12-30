package server

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	fiberCors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/configs"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/logger"
	"time"
)

var AlreadyStartedError = errors.New("server already started")

type Server interface {
	Start(ctx context.Context) error
}

type httpServer struct {
	cfg *configs.Config
	log logger.Logger

	controllers []Controller

	isStarted bool

	app *fiber.App
}

func (s *httpServer) Start(ctx context.Context) error {
	if s.isStarted {
		return AlreadyStartedError
	}

	s.isStarted = true
	defer func() {
		s.isStarted = false
	}()

	go func() {
		for s.isStarted {
			select {
			case <-ctx.Done():
				_ = s.app.ShutdownWithTimeout(time.Minute)
			}
		}
	}()

	if err := s.app.Listen(s.cfg.HttpServerAddr); err != nil {
		return err
	}

	return nil
}

func NewHttpServer(
	cfg *configs.Config,
	log logger.Logger,
	errorHandler fiber.ErrorHandler,
	controllers []Controller,
) Server {
	server := &httpServer{
		cfg:         cfg,
		log:         log,
		controllers: controllers,
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(fiberLogger.New(fiberLogger.Config{
		TimeFormat: time.DateTime,
		Format:     "[${time}] ${status} - ${latency} ${method} ${url}\n",
		Output:     log.Writer(),
	}))

	app.Use(fiberCors.New())
	app.Use(fiberRecover.New())

	for _, controller := range controllers {
		controller.SetupRoutes(app)
	}

	server.app = app

	return server
}
