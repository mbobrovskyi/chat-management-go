package chat

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/aggregate_root"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/user"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/domain/message"
	"time"
)

func NewChat(
	id uint64,
	name string,
	chatType Type,
	image string,
	lastMessage message.Message,
	users []user.User,
	createdBy user.User,
	createdAt time.Time,
	updatedAt time.Time,
) (Chat, error) {
	return &chat{
		AggregateRoot: aggregate_root.New[Chat](id),
		Name:          name,
		Type:          chatType,
		Image:         image,
		LasMessage:    lastMessage,
		Users:         users,
		CreatedBy:     createdBy,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}

func NewDirectChat(
	image string,
	users []user.User,
	createdBy user.User,
) (Chat, error) {
	return NewChat(0, "", Direct, image, nil, users, createdBy, time.Now(), time.Now())
}

func NewGroupChat(
	name string,
	image string,
	users []user.User,
	createdBy user.User,
) (Chat, error) {
	return NewChat(0, name, Group, image, nil, users, createdBy, time.Now(), time.Now())
}
