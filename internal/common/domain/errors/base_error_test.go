package errors

import (
	"fmt"
	"testing"
)

func TestBaseError(t *testing.T) {
	err1 := NewBaseError("Test", "Test", "GetMessage", 500, nil)
	err2 := NewBaseError("Test", "Test", "GetMessage", 500, nil)
	fmt.Println(err1.Equals(err2))
}
