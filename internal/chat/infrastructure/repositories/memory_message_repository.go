package repositories

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/abstracts"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/message"
	"github.com/samber/lo"
)

var _ abstracts.MessageRepository = (*MemoryMessageRepository)(nil)

type MemoryMessageRepository struct {
	messages []message.Message
}

func (m *MemoryMessageRepository) getLastId() uint64 {
	if len(m.messages) == 0 {
		return 0
	}

	found := lo.MaxBy(m.messages, func(a message.Message, b message.Message) bool {
		return a.Id() > b.Id()
	})

	return found.Id()
}

func (m *MemoryMessageRepository) GetAll(ctx context.Context) ([]message.Message, uint64, error) {
	return m.messages, uint64(len(m.messages)), nil
}

func (m *MemoryMessageRepository) Save(ctx context.Context, msg message.Message) (*message.Message, error) {
	var newMessage message.Message

	if msg.Id() == 0 {
		newMessage = message.New(m.getLastId()+1, msg.Text(), msg.Status(), msg.ChatId(),
			msg.CreatedById(), msg.CreatedAt(), msg.UpdatedAt())
	} else {
		newMessage = msg
	}

	m.messages = lo.Filter(m.messages, func(item message.Message, _ int) bool {
		return item.Id() != newMessage.Id()
	})

	m.messages = append(m.messages, newMessage)

	return &newMessage, nil
}

func NewMemoryMessageRepository() *MemoryMessageRepository {
	return &MemoryMessageRepository{}
}
