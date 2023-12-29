package message

import "slices"

type Status uint8

const (
	Draft  Status = 1
	Unread Status = 2
	Read   Status = 3
)

func (ms Status) Uint8() uint8 {
	return uint8(ms)
}

func (ms Status) Types() []Status {
	return []Status{Draft, Unread, Read}
}

func (ms Status) IsValid() bool {
	return slices.Contains(ms.Types(), ms)
}

func NewMessageStatus(messageStatus uint8) Status {
	return Status(messageStatus)
}
