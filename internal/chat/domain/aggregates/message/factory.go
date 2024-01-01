package message

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/valueobjects"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/entities/user"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/errors"
	"strings"
	"time"
)

const (
	minTextSize = 1
	maxTextSize = 5000
)

func New(
	id uint64,
	text string,
	status valueobjects.MessageStatus,
	chatId uint64,
	createdBy uint64,
	createdAt time.Time,
	updatedAt time.Time,
) Message {
	return Message{
		id:          id,
		text:        strings.TrimSpace(text),
		status:      status,
		chatId:      chatId,
		createdById: createdBy,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func NewWithUser(
	id uint64,
	text string,
	status valueobjects.MessageStatus,
	chatId uint64,
	createdBy user.User,
	createdAt time.Time,
	updatedAt time.Time,
) (Message, error) {
	return Message{
		id:          id,
		text:        strings.TrimSpace(text),
		status:      status,
		chatId:      chatId,
		createdById: createdBy.Id(),
		createdBy:   createdBy,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}

func Create(
	text string,
	chatId uint64,
	createdBy user.User,
) (Message, error) {
	text = strings.TrimSpace(text)

	if len(text) < minTextSize {
		return Message{}, errors.NewMinLengthError(minTextSize)
	}

	if len(text) > maxTextSize {
		return Message{}, errors.NewMaxLengthError(maxTextSize)
	}

	return NewWithUser(0, text, valueobjects.Draft, chatId, createdBy, time.Now(), time.Now())
}
