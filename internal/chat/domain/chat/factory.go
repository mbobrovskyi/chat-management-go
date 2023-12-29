package chat

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
	"time"
)

func Create(
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
		Id:         id,
		Name:       name,
		Type:       chatType,
		Image:      image,
		LasMessage: lastMessage,
		CreatedBy:  createdBy,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

func CreateDirect(
	image string,
	createdBy uint64,
) Chat {
	return Create(0, "", Direct, image, nil, createdBy, time.Now(), time.Now())
}

func CreateGroup(
	name string,
	image string,
	createdBy uint64,
) Chat {
	return Create(0, name, Group, image, nil, createdBy, time.Now(), time.Now())
}
