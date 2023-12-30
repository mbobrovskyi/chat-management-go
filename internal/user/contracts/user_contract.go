package contracts

import (
	"github.com/mbobrovskyi/chat-management-go/internal/user/domain"
)

var _ domain.Contract = (*UserContract)(nil)

type UserContract struct{}

func (c *UserContract) GetCurrent(token string) (domain.User, error) {
	// TODO: Get user from management service
	return domain.NewUser(1, "Alice", "Alison", ""), nil
}

func (c *UserContract) GetAll() ([]domain.User, error) {
	// TODO: Search users in management service
	return []domain.User{domain.NewUser(1, "Alice", "Alison", "")}, nil
}

func NewUserContract() *UserContract {
	return &UserContract{}
}
