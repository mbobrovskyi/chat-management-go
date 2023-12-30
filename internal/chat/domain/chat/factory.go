package chat

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
	"time"
)

func Create(
	id uint64,
	name string,
	description string,
	chatType Type,
	image string,
	lastMessage message.Message,
	members []uint64,
	createdBy uint64,
	createdAt time.Time,
	updatedAt time.Time,
) Chat {
	return &chat{
		Id:          id,
		Name:        name,
		Description: description,
		Type:        chatType,
		Image:       image,
		LasMessage:  lastMessage,
		Members:     members,
		CreatedBy:   createdBy,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func CreateDirect(
	image string,
	member uint64,
	createdBy uint64,
) Chat {
	members := []uint64{member, createdBy}
	return Create(0, "", "", Direct, image, nil, members, createdBy, time.Now(), time.Now())
}

func CreateGroup(
	name string,
	image string,
	members []uint64,
	createdBy uint64,
) Chat {
	return Create(0, name, "", Group, image, nil, members, createdBy, time.Now(), time.Now())
}
