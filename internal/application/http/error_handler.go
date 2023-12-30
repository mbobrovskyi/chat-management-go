package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	baseerror2 "github.com/mbobrovskyi/chat-management-go/internal/domain/baseerror"
	configs2 "github.com/mbobrovskyi/chat-management-go/internal/infrastructure/configs"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/logger"
	"maps"
	"net/http"
)

type ErrorHandler struct {
	cfg *configs2.Config
	log logger.Logger
}

func (e *ErrorHandler) Handle(ctx *fiber.Ctx, err error) error {
	var baseError *baseerror2.BaseError

	switch err.(type) {
	case *fiber.Error:
		var fiberErr *fiber.Error
		errors.As(err, &fiberErr)
		switch fiberErr.Code {
		case http.StatusNotFound:
			baseError = baseerror2.NewNotFoundError(err.Error())
		}
	}

	if baseError == nil {
		baseError = baseerror2.NewUndefinedError(err)
	}

	errResponse := &ErrorResponse{
		Timestamp: baseError.GetTimestamp().Format(configs2.DateTimeFormat),
		Code:      baseError.GetCode(),
		Message:   baseError.GetMessage(),
		Metadata:  maps.Clone(baseError.GetMetaData()),
	}

	if baseError.GetHttpStatusCode() >= http.StatusInternalServerError {
		if e.cfg.Environment == configs2.Development {
			errResponse.Metadata["stacktrace"] = baseError.GetStacktrace()
		}
		e.log.Error(baseError.String())
	}

	return ctx.Status(baseError.GetHttpStatusCode()).JSON(errResponse)
}

func NewErrorHandler(cfg *configs2.Config, log logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		cfg: cfg,
		log: log,
	}
}
