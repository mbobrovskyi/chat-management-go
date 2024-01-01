package httpserver

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	fiberCors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mbobrovskyi/chat-management-go/pkg/baseconfig"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/logger"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/startable"
	"time"
)

var AlreadyStartedError = errors.New("server already started")

var _ startable.Startable = (*HttpServer)(nil)

type HttpServer struct {
	cfg *baseconfig.BaseConfig
	log logger.Logger

	controllers []Controller

	isStarted bool

	app *fiber.App
}

func (s *HttpServer) Start(ctx context.Context) error {
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
	cfg *baseconfig.BaseConfig,
	log logger.Logger,
	errorHandler fiber.ErrorHandler,
	controllers []Controller,
) *HttpServer {
	httpServer := &HttpServer{
		cfg:         cfg,
		log:         log,
		controllers: controllers,
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(fiberLogger.New(fiberLogger.Config{
		TimeFormat: baseconfig.DateTimeFormat,
		Format:     "[${time}] ${status} - ${latency} ${method} ${url}\n",
		Output:     log.Writer(),
	}))

	app.Use(fiberCors.New())
	app.Use(fiberRecover.New())

	for _, controller := range controllers {
		controller.SetupRoutes(app)
	}

	httpServer.app = app

	return httpServer
}
