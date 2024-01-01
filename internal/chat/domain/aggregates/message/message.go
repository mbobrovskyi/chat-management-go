package message

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/valueobjects"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/entities/user"
	"time"
)

type Message struct {
	id          uint64
	text        string
	status      valueobjects.MessageStatus
	chatId      uint64
	createdById uint64
	createdBy   user.User
	createdAt   time.Time
	updatedAt   time.Time
}

func (m *Message) Id() uint64 {
	return m.id
}

func (m *Message) Text() string {
	return m.text
}

func (m *Message) Status() valueobjects.MessageStatus {
	return m.status
}

func (m *Message) ChatId() uint64 {
	return m.chatId
}

func (m *Message) CreatedById() uint64 {
	return m.createdById
}

func (m *Message) CreatedBy() user.User {
	return m.createdBy
}

func (m *Message) SetCreatedBy(u user.User) {
	m.createdBy = u
}

func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}

func (m *Message) UpdatedAt() time.Time {
	return m.updatedAt
}
