package domain

import (
	"context"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/message"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/pubsub"
)

type MessageService interface {
	CreateMessage(ctx context.Context, message message.Message) error
}

type messageService struct {
	chatPublisher pubsub.Publisher
}

func (m *messageService) CreateMessage(ctx context.Context, message message.Message) error {
	if err := m.chatPublisher.Publish(ctx, CreateMessagePubSubEventType, 1); err != nil {
		return err
	}
	return nil
}

func NewMessageService(messageRepository MessageRepository, chatPublisher pubsub.Publisher) MessageService {
	return &messageService{chatPublisher: chatPublisher}
}
