package baseerror

import (
	"encoding/json"
	"time"
)

type BaseError struct {
	Timestamp      time.Time      `json:"timestamp"`
	Code           string         `json:"code"`
	Message        string         `json:"message"`
	HttpStatusCode int            `json:"httpStatusCode"`
	Stacktrace     string         `json:"stacktrace"`
	Metadata       map[string]any `json:"metadata"`
}

func (e *BaseError) Error() string {
	return e.Message
}

func (e *BaseError) String() string {
	msg, _ := json.Marshal(e)
	return string(msg)
}

func (e *BaseError) GetTimestamp() time.Time {
	return e.Timestamp
}

func (e *BaseError) GetCode() string {
	return e.Code
}

func (e *BaseError) GetMessage() string {
	return e.Message
}

func (e *BaseError) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e *BaseError) GetStacktrace() string {
	return e.Stacktrace
}

func (e *BaseError) GetMetaData() map[string]any {
	return e.Metadata
}

func (e *BaseError) WithMetadata(key string, value any) *BaseError {
	e.Metadata[key] = value
	return e
}
