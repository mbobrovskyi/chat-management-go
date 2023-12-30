package subscriber

import (
	"context"
	"errors"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/baseerror"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/logger"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/pubsub/common"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"net/http"
	"strconv"
	"strings"
)

var InvalidChannelError = errors.New("invalid channel")

type Subscriber interface {
	Start(ctx context.Context, eventTypes []uint8) error
}

type subscriber struct {
	log          logger.Logger
	rdb          *redis.Client
	prefix       string
	eventHandler EventHandler
}

func (s *subscriber) getEventType(channel string) (uint8, error) {
	parts := strings.Split(channel, s.prefix)
	if len(parts) < 1 {
		return 0, InvalidChannelError
	}

	eventType, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(eventType), nil
}

func (s *subscriber) Start(ctx context.Context, eventTypes []uint8) error {
	channels := lo.Map(eventTypes, func(eventType uint8, _ int) string {
		return common.BuildChannelName(s.prefix, eventType)
	})

	pubsub := s.rdb.Subscribe(ctx, channels...)
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		eventType, err := s.getEventType(msg.Channel)
		if err != nil {
			s.log.Errorf("error on subscriber: %s", err.Error())
		}

		if err := s.eventHandler.Handle(eventType, []byte(msg.Payload)); err != nil {
			s.errorHandle(err)
		}
	}

	return nil
}

func (s *subscriber) errorHandle(err error) {
	baseError := baseerror.NewUndefinedError(err)
	if baseError.HttpStatusCode >= http.StatusInternalServerError {
		s.log.Error(baseError.String())
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
