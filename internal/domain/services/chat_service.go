package services

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/chat"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/interfaces"
)

var _ ChatService = (*ChatServiceImpl)(nil)

type ChatService interface {
	GetById(ctx context.Context, id uint64) (*chat.Chat, error)
	GetAll(ctx context.Context) ([]chat.Chat, uint64, error)
	Save(ctx context.Context, chat chat.Chat) (*chat.Chat, error)
	Delete(ctx context.Context, id uint64) error
}

type ChatServiceImpl struct {
	repository interfaces.ChatRepository
}

func (s *ChatServiceImpl) GetById(ctx context.Context, id uint64) (*chat.Chat, error) {
	return s.repository.GetById(ctx, id)
}

func (s *ChatServiceImpl) GetAll(ctx context.Context) ([]chat.Chat, uint64, error) {
	return s.repository.GetAll(ctx)
}

func (s *ChatServiceImpl) Save(ctx context.Context, chat chat.Chat) (*chat.Chat, error) {
	return s.repository.Save(ctx, chat)
}

func (s *ChatServiceImpl) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, id)
}

func NewChatService(repository interfaces.ChatRepository) *ChatServiceImpl {
	return &ChatServiceImpl{repository: repository}
}
