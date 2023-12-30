package errors

import (
	"fmt"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/baseerror"
)

func NewValueIsRequiredError() error {
	return baseerror.NewValidationError("Value is required.")
}

func NewMinLengthError(minLength int) error {
	return baseerror.NewValidationError(fmt.Sprintf("Min length is %d.", minLength))
}

func NewMaxLengthError(maxLength int) error {
	return baseerror.NewValidationError(fmt.Sprintf("Max length is %d.", maxLength))
}

func NewValueIsNotValidError() error {
	return baseerror.NewValidationError("Value is not valid.")
}
