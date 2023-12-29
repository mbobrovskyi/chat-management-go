package repositories

import "github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain"

var _ domain.MessageRepository = (*MessageRepository)(nil)

type MessageRepository struct{}

func NewMessageRepository() domain.MessageRepository {
	return MessageRepository{}
}
