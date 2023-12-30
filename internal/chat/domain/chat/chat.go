package chat

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
	"time"
)

type Chat interface {
	GetId() uint64
	GetName() string
	GetDescription() string
	GetType() Type
	GetImage() string
	GetLastMessage() message.Message
	GetMembers() []uint64
	GetCreatedBy() uint64
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type chat struct {
	Id          uint64
	Name        string
	Description string
	Type        Type
	Image       string
	LasMessage  message.Message
	Members     []uint64
	CreatedBy   uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c *chat) GetId() uint64 {
	return c.Id
}

func (c *chat) GetName() string {
	return c.Name
}

func (c *chat) GetDescription() string {
	return c.Description
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

func (c *chat) GetMembers() []uint64 {
	return c.Members
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
