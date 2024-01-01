package common

import (
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/baseerror"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/logger"
	"net/http"
)

type BaseErrorHandler struct {
	log logger.Logger
}

func (eh *BaseErrorHandler) Handle(err error) {
	baseError := baseerror.NewUndefinedError(err)
	if baseError.HttpStatusCode >= http.StatusInternalServerError {
		eh.log.Error(baseError.String())
	}
}

func NewBaseErrorHandler(log logger.Logger) *BaseErrorHandler {
	return &BaseErrorHandler{log: log}
}
