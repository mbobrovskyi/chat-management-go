package repositories

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
	"github.com/samber/lo"
)

type MessageRepository struct {
	messages []message.Message
}

func (m *MessageRepository) getLastId() uint64 {
	if len(m.messages) == 0 {
		return 0
	}

	return lo.MaxBy(m.messages, func(a message.Message, b message.Message) bool {
		return a.GetId() > b.GetId()
	}).GetId()
}

func (m *MessageRepository) GetAll(ctx context.Context) ([]message.Message, uint64, error) {
	return m.messages, uint64(len(m.messages)), nil
}

func (m *MessageRepository) Save(ctx context.Context, msg message.Message) (message.Message, error) {
	var newMessage message.Message

	if msg.GetId() == 0 {
		newMessage = message.Create(m.getLastId()+1, msg.GetText(), msg.GetStatus(), msg.GetChatId(),
			msg.GetCreatedBy(), msg.GetCreatedAt(), msg.GetUpdatedAt())
	} else {
		newMessage = msg
	}

	m.messages = lo.Filter(m.messages, func(item message.Message, _ int) bool {
		return item.GetId() != newMessage.GetId()
	})

	m.messages = append(m.messages, newMessage)

	return newMessage, nil
}

func NewMessageRepository() message.Repository {
	return &MessageRepository{}
}
