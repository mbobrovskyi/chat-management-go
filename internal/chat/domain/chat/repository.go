package chat

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Chat, uint64, error)
	GetById(ctx context.Context, id uint64) (Chat, error)
	Save(ctx context.Context, chat Chat) (Chat, error)
	Delete(ctx context.Context, id uint64) error
}
