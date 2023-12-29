package chat

import (
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]Chat, uint64, error)
	GetById(ctx context.Context, id uint64) (Chat, error)
	Save(ctx context.Context, chat Chat) (Chat, error)
	Delete(ctx context.Context, id uint64) error
}

type chatService struct {
	repository Repository
}

func (s *chatService) GetAll(ctx context.Context) ([]Chat, uint64, error) {
	return s.repository.GetAll(ctx)
}

func (s *chatService) GetById(ctx context.Context, id uint64) (Chat, error) {
	return s.repository.GetById(ctx, id)
}

func (s *chatService) Save(ctx context.Context, chat Chat) (Chat, error) {
	return s.repository.Save(ctx, chat)
}

func (s *chatService) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, id)
}

func NewService(repository Repository) Service {
	return &chatService{repository: repository}
}
