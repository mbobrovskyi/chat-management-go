package domain

import (
	"context"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/chat"
)

type ChatRepository interface {
	GetAll(ctx context.Context) ([]chat.Chat, error)
	GetById(ctx context.Context, id uint64) (chat.Chat, error)
	Save(ctx context.Context, chat chat.Chat) (chat.Chat, error)
	Delete(ctx context.Context, id uint64) error
}
