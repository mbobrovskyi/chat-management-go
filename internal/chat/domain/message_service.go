package domain

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/common"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/publisher"
)

type MessageService interface {
	CreateMessage(ctx context.Context, message message.Message) error
}

type messageService struct {
	chatPublisher publisher.Publisher
}

func (m *messageService) CreateMessage(ctx context.Context, message message.Message) error {
	if err := m.chatPublisher.Publish(ctx, CreateMessagePubSubEventType, common.MessageToDTO(message)); err != nil {
		return err
	}

	return nil
}

func NewMessageService(messageRepository MessageRepository, chatPublisher publisher.Publisher) MessageService {
	return &messageService{chatPublisher: chatPublisher}
}
