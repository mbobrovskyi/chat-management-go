package chat

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/message"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/aggregate_root"
	"time"
)

type Chat interface {
	aggregate_root.AggregateRoot[Chat]

	GetName() string
	GetType() Type
	GetImage() string
	GetLastMessage() message.Message
	GetCreatedBy() uint64
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type chat struct {
	aggregate_root.AggregateRoot[Chat]

	Name       string
	Type       Type
	Image      string
	LasMessage message.Message
	CreatedBy  uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (c *chat) GetName() string {
	return c.Name
}

func (c *chat) GetType() Type {
	return c.Type
}

func (c *chat) GetImage() string {
	return c.Image
}

func (c *chat) GetLastMessage() message.Message {
	return c.LasMessage
}

func (c *chat) GetCreatedBy() uint64 {
	return c.CreatedBy
}

func (c *chat) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *chat) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}
