package session

import "github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/user"

type Session interface {
	GetUser() user.User
}
