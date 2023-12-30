package baseerror

import (
	"net/http"
	"runtime/debug"
	"time"
)

func NewBaseError(
	code string,
	message string,
	httpStatusCode int,
) *BaseError {
	return &BaseError{
		Timestamp:      time.Now(),
		Code:           code,
		Message:        message,
		HttpStatusCode: httpStatusCode,
		Stacktrace:     string(debug.Stack()),
		Metadata:       make(map[string]any),
	}
}

func NewValidationError(message string) *BaseError {
	return NewBaseError("ValidationError", message, http.StatusBadRequest)
}

func NewUnauthorizedError(message string) *BaseError {
	return NewBaseError("UnauthorizedError", message, http.StatusUnauthorized)
}

func NewNotFoundError(message string) *BaseError {
	return NewBaseError("NotFoundError", message, http.StatusNotFound)
}

func NewConflictError(message string) *BaseError {
	return NewBaseError("ConflictError", message, http.StatusConflict)
}

func NewUndefinedError(err error) *BaseError {
	baseError, ok := err.(*BaseError)
	if !ok {
		baseError = NewBaseError("UndefinedError", err.Error(), http.StatusInternalServerError)
	}

	return baseError
}
