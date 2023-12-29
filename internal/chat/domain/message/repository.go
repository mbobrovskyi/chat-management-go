package message

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Message, uint64, error)
	Save(ctx context.Context, msg Message) (Message, error)
}
