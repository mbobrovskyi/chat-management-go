package publisher

import (
	"context"
	"encoding/json"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub"
	"github.com/redis/go-redis/v9"
)

var _ Publisher = (*RedisPublisher)(nil)

type RedisPublisher struct {
	redisClient *redis.Client
	prefix      string
}

func (ps *RedisPublisher) Publish(ctx context.Context, eventType int, data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := ps.redisClient.Publish(ctx, pubsub.BuildChannelName(ps.prefix, eventType), bytes).Err(); err != nil {
		return err
	}

	return nil
}

func NewRedisPublisher(redisClient *redis.Client, prefix string) *RedisPublisher {
	return &RedisPublisher{
		redisClient: redisClient,
		prefix:      prefix,
	}
}
