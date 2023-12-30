package chat

import (
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/message"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/valueobjects"
	"time"
)

type Chat struct {
	id          uint64
	name        string
	description string
	chatType    valueobjects.ChatType
	image       string
	lastMessage *message.Message
	memberIds   []uint64
	createdBy   uint64
	createdAt   time.Time
	updatedAt   time.Time
}

func (c *Chat) Id() uint64 {
	return c.id
}

func (c *Chat) Name() string {
	return c.name
}

func (c *Chat) Description() string {
	return c.description
}

func (c *Chat) Type() valueobjects.ChatType {
	return c.chatType
}

func (c *Chat) Image() string {
	return c.image
}

func (c *Chat) LastMessage() *message.Message {
	return c.lastMessage
}

func (c *Chat) MemberIds() []uint64 {
	return c.memberIds
}

func (c *Chat) CreatedBy() uint64 {
	return c.createdBy
}

func (c *Chat) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Chat) UpdatedAt() time.Time {
	return c.updatedAt
}
