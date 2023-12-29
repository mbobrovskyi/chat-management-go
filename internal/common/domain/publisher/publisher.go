package publisher

import (
	"context"
	"encoding/json"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/common"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/http_error"
	"github.com/redis/go-redis/v9"
)

type Publisher interface {
	Publish(ctx context.Context, eventType uint8, data any) error
}

type publisher struct {
	rdb    *redis.Client
	prefix string
}

func (ps *publisher) Publish(ctx context.Context, eventType uint8, data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := ps.rdb.Publish(ctx, common.BuildChannelName(ps.prefix, eventType), bytes).Err(); err != nil {
		return http_error.NewPublisherError(err.Error())
	}

	return nil
}

func NewPublisher(rdb *redis.Client, prefix string) Publisher {
	return &publisher{
		rdb:    rdb,
		prefix: prefix,
	}
}
