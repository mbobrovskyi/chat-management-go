package publisher

import (
	"context"
	"encoding/json"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub/event"
)

var _ Publisher = (*MemoryPublisher)(nil)

type MemoryPublisher struct {
	ch     chan event.Event
	prefix string
}

func (ps *MemoryPublisher) Publish(ctx context.Context, eventType int, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ps.ch <- event.New(eventType, jsonData)

	return nil
}

func NewMemoryPublisher(ch chan event.Event) *MemoryPublisher {
	return &MemoryPublisher{ch: ch}
}
