package repositories

import (
	"context"
	message2 "github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/message"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/interfaces"
	"github.com/samber/lo"
)

var _ interfaces.MessageRepository = (*MemoryMessageRepository)(nil)

type MemoryMessageRepository struct {
	messages []message2.Message
}

func (m *MemoryMessageRepository) getLastId() uint64 {
	if len(m.messages) == 0 {
		return 0
	}

	found := lo.MaxBy(m.messages, func(a message2.Message, b message2.Message) bool {
		return a.Id() > b.Id()
	})

	return found.Id()
}

func (m *MemoryMessageRepository) GetAll(ctx context.Context) ([]message2.Message, uint64, error) {
	return m.messages, uint64(len(m.messages)), nil
}

func (m *MemoryMessageRepository) Save(ctx context.Context, msg message2.Message) (*message2.Message, error) {
	var newMessage message2.Message

	if msg.Id() == 0 {
		newMessage = message2.New(m.getLastId()+1, msg.Text(), msg.Status(), msg.ChatId(),
			msg.CreatedById(), msg.CreatedAt(), msg.UpdatedAt())
	} else {
		newMessage = msg
	}

	m.messages = lo.Filter(m.messages, func(item message2.Message, _ int) bool {
		return item.Id() != newMessage.Id()
	})

	m.messages = append(m.messages, newMessage)

	return &newMessage, nil
}

func NewMemoryMessageRepository() *MemoryMessageRepository {
	return &MemoryMessageRepository{}
}
