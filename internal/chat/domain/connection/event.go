package connection

import (
	"encoding/json"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/value_object"
)

type Event interface {
	value_object.ValueObject[*event]

	GetType() uint8
	GetData() json.RawMessage
}

type event struct {
	Type uint8
	Data json.RawMessage
}

func (e *event) Equals(other *event) bool {
	if other == nil {
		return false
	}

	data, _ := e.GetData().MarshalJSON()
	dataOther, _ := other.GetData().MarshalJSON()

	return e.Type == other.GetType() && string(data) == string(dataOther)
}

func (e *event) GetType() uint8 {
	return e.Type
}

func (e *event) GetData() json.RawMessage {
	return e.Data
}

func NewEvent(eventType uint8, data json.RawMessage) Event {
	return &event{
		Type: eventType,
		Data: data,
	}
}
