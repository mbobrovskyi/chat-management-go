package interfaces

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/message"
)

type MessageRepository interface {
	GetAll(ctx context.Context) ([]message.Message, uint64, error)
	Save(ctx context.Context, msg message.Message) (*message.Message, error)
}
