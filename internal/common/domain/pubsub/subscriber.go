package pubsub

import (
	"context"
	"errors"
	"fmt"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/infrastructure/logger"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

var InvalidChannelError = errors.New("invalid channel")

type Subscriber interface {
	Subscribe(ctx context.Context, eventTypes []uint8) error
}

type subscriber struct {
	log          logger.Logger
	rdb          *redis.Client
	prefix       string
	eventHandler EventHandler
}

func (ps *subscriber) getEventType(channel string) (uint8, error) {
	parts := strings.Split(channel, ps.prefix)
	if len(parts) < 1 {
		return 0, InvalidChannelError
	}

	eventType, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(eventType), nil
}

func (ps *subscriber) Subscribe(ctx context.Context, eventTypes []uint8) error {
	channels := lo.Map(eventTypes, func(eventType uint8, _ int) string {
		return fmt.Sprintf("%s%d", ps.prefix, eventType)
	})

	pubsub := ps.rdb.Subscribe(ctx, channels...)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			return fmt.Errorf("error on pubsub: %w", err)
		}

		eventType, err := ps.getEventType(msg.Channel)
		if err != nil {
			ps.log.Errorf("error on pub/sub: %s", err.Error())
		}

		if err := ps.eventHandler.Handle(eventType, []byte(msg.Payload)); err != nil {
			ps.log.Errorf("error on pub/sub: %s", err.Error())
		}
	}
}

func NewSubscriber(log logger.Logger, rdb *redis.Client, eventHandler EventHandler, prefix string) Subscriber {
	return &subscriber{
		log:          log,
		rdb:          rdb,
		eventHandler: eventHandler,
		prefix:       prefix,
	}
}
