package abstracts

import (
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/entities/user"
)

type UserClient interface {
	GetCurrent(token string) (*user.User, error)
	GetAll(token string) ([]user.User, error)
}
