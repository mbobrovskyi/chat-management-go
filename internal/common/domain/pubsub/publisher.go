package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/http_error"
	"github.com/redis/go-redis/v9"
)

type Publisher interface {
	Publish(ctx context.Context, eventType uint8, data any) error
}

type publisher struct {
	rdb    *redis.Client
	prefix string
}

func (ps *publisher) buildChannel(eventType uint8) string {
	return fmt.Sprintf("%s%d", ps.prefix, eventType)
}

func (ps *publisher) Publish(ctx context.Context, eventType uint8, data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return http_error.NewPubSubError(err.Error())
	}

	if err := ps.rdb.Publish(ctx, ps.buildChannel(eventType), bytes).Err(); err != nil {
		return http_error.NewPubSubError(err.Error())
	}

	return nil
}

func NewPublisher(rdb *redis.Client, prefix string) Publisher {
	return &publisher{
		rdb:    rdb,
		prefix: prefix,
	}
}
