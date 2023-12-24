package errors

import (
	"encoding/json"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/value_object"
	"runtime/debug"
)

type BaseError interface {
	error
	value_object.ValueObject[BaseError]

	String() string
	GetDomain() string
	GetMessage() string
	GetStack() string
	GetCode() string
	GetHttpStatusCode() int
	GetMetaData() map[string]any
}

type errorData struct {
	Domain         string         `json:"domain"`
	Code           string         `json:"code"`
	Message        string         `json:"message"`
	Stack          string         `json:"-"`
	HttpStatusCode int            `json:"httpStatusCode"`
	MetaData       map[string]any `json:"metaData"`
}

func (e *errorData) Error() string {
	return e.String()
}

func (e *errorData) Equals(other BaseError) bool {
	if other == nil {
		return false
	}

	if e == other {
		return true
	}

	return e.GetDomain() == other.GetDomain() && e.GetCode() == other.GetCode()
}

func (e *errorData) String() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func (e *errorData) GetDomain() string {
	return e.Domain
}

func (e *errorData) GetCode() string {
	return e.Code
}

func (e *errorData) GetMessage() string {
	return e.Message
}

func (e *errorData) GetStack() string {
	return e.Stack
}

func (e *errorData) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e *errorData) GetMetaData() map[string]any {
	return e.MetaData
}

func NewBaseErrorWithMessage(
	domain string,
	code string,
	message string,
	httpStatusCode int,
) BaseError {
	return &errorData{
		Domain:         domain,
		Code:           code,
		Message:        message,
		HttpStatusCode: httpStatusCode,
	}
}

func NewBaseErrorWithError(
	domain string,
	code string,
	err error,
	httpStatusCode int,
) BaseError {
	var stackStr string
	if err != nil {
		stackStr = string(debug.Stack())
	}

	return &errorData{
		Domain:         domain,
		Code:           code,
		Message:        err.Error(),
		Stack:          stackStr,
		HttpStatusCode: httpStatusCode,
	}
}
