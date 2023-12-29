package message

import (
	"strings"
	"time"
)

func Create(
	id uint64,
	text string,
	status Status,
	chatId uint64,
	createdBy uint64,
	createdAt time.Time,
	updatedAt time.Time,
) *message {
	return &message{
		id:        id,
		text:      strings.TrimSpace(text),
		status:    status,
		chatId:    chatId,
		createdBy: createdBy,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func CreateNew(
	text string,
	chatId uint64,
	createdBy uint64,
) Message {
	return Create(0, text, Draft, chatId, createdBy, time.Now(), time.Now())
}
