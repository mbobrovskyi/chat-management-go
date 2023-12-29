package repositories

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
)

type MessageRepository struct{}

func (m *MessageRepository) GetMessages() ([]message.Message, uint64, error) {
	return nil, 0, nil
}

func NewMessageRepository() message.Repository {
	return &MessageRepository{}
}
