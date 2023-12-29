package http_error

import (
	"encoding/json"
	"runtime/debug"
	"time"
)

type HttpError interface {
	error

	String() string
	GetTimestamp() time.Time
	GetCode() string
	GetMessage() string
	GetHttpStatusCode() int
	GetStacktrace() string
	GetMetaData() map[string]any
	WithMetadata(key string, value any) HttpError
}

type httpError struct {
	Timestamp      time.Time
	Code           string
	Message        string
	HttpStatusCode int
	Stacktrace     string
	Metadata       map[string]any
}

func (e *httpError) Error() string {
	return e.String()
}

func (e *httpError) String() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func (e *httpError) GetTimestamp() time.Time {
	return e.Timestamp
}

func (e *httpError) GetCode() string {
	return e.Code
}

func (e *httpError) GetMessage() string {
	return e.Message
}

func (e *httpError) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e *httpError) GetStacktrace() string {
	return e.Stacktrace
}

func (e *httpError) GetMetaData() map[string]any {
	return e.Metadata
}

func (e *httpError) WithMetadata(key string, value any) HttpError {
	e.Metadata[key] = value
	return e
}

func NewHttpError(
	code string,
	message string,
	httpStatusCode int,
) HttpError {
	return &httpError{
		Timestamp:      time.Now(),
		Code:           code,
		Message:        message,
		HttpStatusCode: httpStatusCode,
		Stacktrace:     string(debug.Stack()),
		Metadata:       make(map[string]any),
	}
}
