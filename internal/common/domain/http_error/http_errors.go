package http_error

import "net/http"

func NewNotFoundError(message string) HttpError {
	return NewHttpError("NotFoundError", message, http.StatusBadRequest)
}
