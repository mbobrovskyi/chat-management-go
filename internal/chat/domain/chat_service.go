package domain

import (
	"context"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/chat"
)

type ChatService interface {
	GetAll(ctx context.Context) ([]chat.Chat, error)
	GetById(ctx context.Context, id uint64) ([]chat.Chat, error)
	Save(ctx context.Context, chat chat.Chat) (chat.Chat, error)
	Delete(ctx context.Context, id uint64) error
}

type chatService struct {
	repository ChatRepository
}

func (s *chatService) GetAll(ctx context.Context) ([]chat.Chat, error) {
	return s.repository.GetAll(ctx)
}

func (s *chatService) GetById(ctx context.Context, id uint64) ([]chat.Chat, error) {
	return s.repository.GetAll(ctx)
}

func (s *chatService) Save(ctx context.Context, chat chat.Chat) (chat.Chat, error) {
	return s.repository.Save(ctx, chat)
}

func (s *chatService) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, id)
}

func NewService(repository ChatRepository) ChatService {
	return &chatService{repository: repository}
}
