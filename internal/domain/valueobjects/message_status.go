package valueobjects

import (
	"github.com/mbobrovskyi/chat-management-go/internal/domain/errors"
	"slices"
)

type MessageStatus uint8

const (
	Draft  MessageStatus = 1
	Unread MessageStatus = 2
	Read   MessageStatus = 3
)

func (s MessageStatus) Uint8() uint8 {
	return uint8(s)
}

func Statuses() []MessageStatus {
	return []MessageStatus{Draft, Unread, Read}
}

func NewMessageStatus(messageStatus uint8) (MessageStatus, error) {
	if !slices.Contains(Statuses(), MessageStatus(messageStatus)) {
		return Draft, errors.NewValueIsNotValidError()
	}

	return MessageStatus(messageStatus), nil
}
