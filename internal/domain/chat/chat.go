package chat

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/aggregate_root"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/user"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/domain/message"
	"time"
)

type Chat interface {
	aggregate_root.AggregateRoot[Chat]
}

type chat struct {
	aggregate_root.AggregateRoot[Chat]

	Name       string
	Type       Type
	Image      string
	LasMessage message.Message
	Users      []user.User
	CreatedBy  user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
