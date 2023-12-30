package session

import (
	"github.com/google/uuid"
	"github.com/mbobrovskyi/chat-management-go/internal/user/domain"
)

type Session interface {
	GetId() uuid.UUID
	GetUser() domain.User
}

type session struct {
	Id   uuid.UUID
	User domain.User
}

func (s *session) GetId() uuid.UUID {
	return s.Id
}

func (s *session) GetUser() domain.User {
	return s.User
}

func NewSession(user domain.User) Session {
	return &session{
		Id:   uuid.New(),
		User: user,
	}
}
