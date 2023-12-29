package http_error

import "net/http"

func NewNotFoundError(message string) HttpError {
	return NewHttpError("NotFoundError", message, http.StatusNotFound)
}

func NewPubSubError(message string) HttpError {
	return NewHttpError("PubSubError", message, http.StatusInternalServerError)
}
