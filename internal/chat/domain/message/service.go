package message

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/publisher"
)

type Service interface {
	CreateMessage(ctx context.Context, message Message) error
}

type messageService struct {
	messageRepository Repository
	chatPublisher     publisher.Publisher
}

func (m *messageService) CreateMessage(ctx context.Context, message Message) error {
	// TODO: Create on database
	if err := m.chatPublisher.Publish(ctx, domain.CreateMessagePubSubEventType, toDTO(message)); err != nil {
		return err
	}
	return nil
}

func NewMessageService(
	messageRepository Repository,
	chatPublisher publisher.Publisher,
) Service {
	return &messageService{
		messageRepository: messageRepository,
		chatPublisher:     chatPublisher,
	}
}
