package chat

import (
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/message"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/entities/user"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/valueobjects"
	"time"
)

func New(
	id uint64,
	name string,
	description string,
	chatType valueobjects.ChatType,
	image string,
	lastMessage *message.Message,
	members []uint64,
	createdBy uint64,
	createdAt time.Time,
	updatedAt time.Time,
) Chat {
	return Chat{
		id:          id,
		name:        name,
		description: description,
		chatType:    chatType,
		image:       image,
		lastMessage: lastMessage,
		memberIds:   members,
		createdBy:   createdBy,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func CreateDirect(
	image string,
	member uint64,
	createdBy user.User,
) (Chat, error) {
	// TODO: Check is valid member

	return Chat{
		chatType:  valueobjects.Direct,
		image:     image,
		memberIds: []uint64{member, createdBy.Id()},
		createdBy: createdBy.Id(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func CreateGroup(
	name string,
	image string,
	createdBy user.User,
) (Chat, error) {

	return Chat{
		name:      name,
		chatType:  valueobjects.Group,
		image:     image,
		memberIds: []uint64{createdBy.Id()},
		createdBy: createdBy.Id(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}
