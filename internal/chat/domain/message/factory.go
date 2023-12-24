package message

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/aggregate_root"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/user"
	"strings"
	"time"
)

func NewMessage(
	id uint64,
	text string,
	status Status,
	chatId uint64,
	createdBy user.User,
	createdAt time.Time,
	updatedAt time.Time,
) (*message, error) {
	text = strings.TrimSpace(text)
	return &message{
		AggregateRoot: aggregate_root.New[Message](id),
		Text:          text,
		Status:        status,
		ChatId:        chatId,
		CreatedBy:     createdBy,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}

func NewDefaultMessage(
	text string,
	chatId uint64,
	createdBy user.User,
) (Message, error) {
	return NewMessage(0, text, Draft, chatId, createdBy, time.Now(), time.Now())
}
