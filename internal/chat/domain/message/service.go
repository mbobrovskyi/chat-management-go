package message

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/publisher"
)

type Service interface {
	GetAll(ctx context.Context) ([]Message, uint64, error)
	Create(ctx context.Context, message Message) error
}

type messageService struct {
	messageRepository Repository
	chatPublisher     publisher.Publisher
}

func (m *messageService) GetAll(ctx context.Context) ([]Message, uint64, error) {
	return m.messageRepository.GetAll(ctx)
}

func (m *messageService) Create(ctx context.Context, msg Message) error {
	newMsg, err := m.messageRepository.Save(ctx, msg)
	if err != nil {
		return err
	}

	if err := m.chatPublisher.Publish(ctx, domain.CreateMessagePubSubEventType, toDTO(newMsg)); err != nil {
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
