package contracts

import (
	"github.com/mbobrovskyi/chat-management-go/internal/domain/entities/user"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/interfaces"
)

var _ interfaces.UserContract = (*MockUserContract)(nil)

type MockUserContract struct{}

func (c *MockUserContract) GetCurrent(token string) (*user.User, error) {
	user := user.New(1, "Alice", "Alison", "")
	return &user, nil
}

func (c *MockUserContract) GetAll(token string) ([]user.User, error) {
	return []user.User{
		user.New(1, "Alice", "Alison", ""),
		user.New(2, "Bob", "Bobson", ""),
	}, nil
}

func NewUserContract() *MockUserContract {
	return &MockUserContract{}
}
