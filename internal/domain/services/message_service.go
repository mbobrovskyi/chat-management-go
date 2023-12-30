package services

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/message"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/interfaces"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/pubsub/publisher"
)

var _ MessageService = (*MessageServiceImpl)(nil)

type MessageService interface {
	GetAll(ctx context.Context) ([]message.Message, uint64, error)
	Save(ctx context.Context, message message.Message) (*message.Message, error)
}

type MessageServiceImpl struct {
	messageRepository interfaces.MessageRepository
	chatPublisher     publisher.Publisher
}

func (m *MessageServiceImpl) GetAll(ctx context.Context) ([]message.Message, uint64, error) {
	return m.messageRepository.GetAll(ctx)
}

func (m *MessageServiceImpl) Save(ctx context.Context, msg message.Message) (*message.Message, error) {
	newMsg, err := m.messageRepository.Save(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := m.chatPublisher.Publish(ctx, domain.CreateMessagePubSubEventType, MessageToDTO(*newMsg)); err != nil {
		return nil, err
	}

	return newMsg, nil
}

func NewMessageService(
	messageRepository interfaces.MessageRepository,
	chatPublisher publisher.Publisher,
) *MessageServiceImpl {
	return &MessageServiceImpl{
		messageRepository: messageRepository,
		chatPublisher:     chatPublisher,
	}
}
