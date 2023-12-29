package http_error

import (
	"net/http"
)

func NewNotFoundError(message string) HttpError {
	return NewHttpError("NotFoundError", message, http.StatusNotFound)
}

func NewUnauthorizedError(message string) HttpError {
	return NewHttpError("UnauthorizedError", message, http.StatusUnauthorized)
}

func NewPublisherError(message string) HttpError {
	return NewHttpError("PublisherError", message, http.StatusInternalServerError)
}
