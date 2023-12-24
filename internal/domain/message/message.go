package message

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/aggregate_root"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/user"
	"time"
)

type Message interface {
	aggregate_root.AggregateRoot[Message]
}

type message struct {
	aggregate_root.AggregateRoot[Message]

	Text      string
	Status    Status
	ChatId    uint64
	CreatedBy user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
