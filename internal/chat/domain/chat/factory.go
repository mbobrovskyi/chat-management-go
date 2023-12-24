package chat

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/message"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/aggregate_root"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/user"
	"time"
)

func NewChat(
	id uint64,
	name string,
	chatType Type,
	image string,
	lastMessage message.Message,
	createdBy uint64,
	createdAt time.Time,
	updatedAt time.Time,
) Chat {
	return &chat{
		AggregateRoot: aggregate_root.New[Chat](id),
		Name:          name,
		Type:          chatType,
		Image:         image,
		LasMessage:    lastMessage,
		CreatedBy:     createdBy,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func NewDirectChat(
	image string,
	users []user.User,
	createdBy uint64,
) Chat {
	return NewChat(0, "", Direct, image, nil, createdBy, time.Now(), time.Now())
}

func NewGroupChat(
	name string,
	image string,
	users []user.User,
	createdBy uint64,
) Chat {
	return NewChat(0, name, Group, image, nil, createdBy, time.Now(), time.Now())
}
