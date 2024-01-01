package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mbobrovskyi/chat-management-go/config"
	"github.com/mbobrovskyi/chat-management-go/pkg/baseconfig"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/baseerror"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/logger"
	"maps"
	"net/http"
)

type ErrorHandler struct {
	cfg *config.Config
	log logger.Logger
}

func (e *ErrorHandler) Handle(ctx *fiber.Ctx, err error) error {
	var baseError *baseerror.BaseError

	switch err.(type) {
	case *fiber.Error:
		var fiberErr *fiber.Error
		errors.As(err, &fiberErr)
		switch fiberErr.Code {
		case http.StatusNotFound:
			baseError = baseerror.NewNotFoundError(err.Error())
		}
	}

	if baseError == nil {
		baseError = baseerror.NewUndefinedError(err)
	}

	errResponse := &ErrorResponse{
		Timestamp: baseError.GetTimestamp().Format(baseconfig.DateTimeFormat),
		Code:      baseError.GetCode(),
		Message:   baseError.GetMessage(),
		Metadata:  maps.Clone(baseError.GetMetaData()),
	}

	if baseError.GetHttpStatusCode() >= http.StatusInternalServerError {
		if e.cfg.Environment == baseconfig.Development {
			errResponse.Metadata["stacktrace"] = baseError.GetStacktrace()
		}
		e.log.Error(baseError.String())
	}

	return ctx.Status(baseError.GetHttpStatusCode()).JSON(errResponse)
}

func NewErrorHandler(cfg *config.Config, log logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		cfg: cfg,
		log: log,
	}
}
