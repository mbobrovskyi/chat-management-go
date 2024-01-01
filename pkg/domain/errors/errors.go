package errors

import (
	"fmt"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/baseerror"
)

func NewValueIsRequiredError() *baseerror.BaseError {
	return baseerror.NewValidationError("Value is required.")
}

func NewMinLengthError(minLength int) *baseerror.BaseError {
	return baseerror.NewValidationError(fmt.Sprintf("Min length is %d.", minLength))
}

func NewMaxLengthError(maxLength int) *baseerror.BaseError {
	return baseerror.NewValidationError(fmt.Sprintf("Max length is %d.", maxLength))
}

func NewValueIsNotValidError() *baseerror.BaseError {
	return baseerror.NewValidationError("Value is not valid.")
}
