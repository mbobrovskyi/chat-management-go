package interfaces

import (
	"github.com/mbobrovskyi/chat-management-go/internal/domain/entities/user"
)

type UserContract interface {
	GetCurrent(token string) (*user.User, error)
	GetAll(token string) ([]user.User, error)
}
