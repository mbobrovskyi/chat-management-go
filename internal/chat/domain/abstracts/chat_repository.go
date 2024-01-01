package abstracts

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/chat"
)

type ChatRepository interface {
	GetAll(ctx context.Context) ([]chat.Chat, uint64, error)
	GetById(ctx context.Context, id uint64) (*chat.Chat, error)
	Save(ctx context.Context, chat chat.Chat) (*chat.Chat, error)
	Delete(ctx context.Context, id uint64) error
}
