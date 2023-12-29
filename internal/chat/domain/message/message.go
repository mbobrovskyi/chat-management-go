package message

import (
	"time"
)

type Message interface {
	GetId() uint64
	GetText() string
	GetStatus() Status
	GetChatId() uint64
	GetCreatedBy() uint64
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type message struct {
	id        uint64
	text      string
	status    Status
	chatId    uint64
	createdBy uint64
	createdAt time.Time
	updatedAt time.Time
}

func (m *message) GetId() uint64 {
	return m.id
}

func (m *message) GetText() string {
	return m.text
}

func (m *message) GetStatus() Status {
	return m.status
}

func (m *message) GetChatId() uint64 {
	return m.chatId
}

func (m *message) GetCreatedBy() uint64 {
	return m.createdBy
}

func (m *message) GetCreatedAt() time.Time {
	return m.createdAt
}

func (m *message) GetUpdatedAt() time.Time {
	return m.updatedAt
}
