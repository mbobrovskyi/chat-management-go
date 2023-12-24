package domain

import "github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/chat"

type Repository interface {
	GetAll() ([]chat.Chat, error)
	GetById(id uint64) (chat.Chat, error)
	Save(chat chat.Chat) (chat.Chat, error)
	Delete(id uint64) error
}
