package valueobjects

import (
	"github.com/mbobrovskyi/chat-management-go/internal/domain/errors"
	"slices"
)

type ChatType uint8

const (
	Direct ChatType = 1
	Group  ChatType = 2
)

func (t ChatType) Uint8() uint8 {
	return uint8(t)
}

func Types() []ChatType {
	return []ChatType{Direct, Group}
}

func NewType(chatType uint8) (ChatType, error) {
	if !slices.Contains(Types(), ChatType(chatType)) {
		return Direct, errors.NewValueIsNotValidError()
	}
	return ChatType(chatType), nil
}
