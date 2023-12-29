package session

import (
	"github.com/google/uuid"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/user"
)

type Session interface {
	GetId() uuid.UUID
	GetUser() user.User
}

type session struct {
	Id   uuid.UUID
	User user.User
}

func (s *session) GetId() uuid.UUID {
	return s.Id
}

func (s *session) GetUser() user.User {
	return s.User
}

func NewSession(user user.User) Session {
	return &session{
		Id:   uuid.New(),
		User: user,
	}
}
