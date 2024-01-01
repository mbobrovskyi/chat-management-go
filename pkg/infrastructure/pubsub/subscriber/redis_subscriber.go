package subscriber

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/baseerror"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/logger"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/startable"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"net/http"
	"strconv"
	"strings"
)

var _ startable.Startable = (*RedisSubscriber)(nil)

type RedisSubscriber struct {
	log          logger.Logger
	redisClient  *redis.Client
	prefix       string
	eventTypes   []int
	eventHandler EventHandler
	isStarted    bool
}

func (s *RedisSubscriber) Start(ctx context.Context) error {
	if s.isStarted {
		return AlreadyStartedError
	}

	s.isStarted = true
	defer func() {
		s.isStarted = false
	}()

	channels := lo.Map(s.eventTypes, func(eventType int, _ int) string {
		return pubsub.BuildChannelName(s.prefix, eventType)
	})

	pubsub := s.redisClient.Subscribe(ctx, channels...)
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		eventType, err := s.getEventType(msg.Channel)
		if err != nil {
			return err
		}

		if err := s.eventHandler.Handle(eventType, []byte(msg.Payload)); err != nil {
			s.errorHandle(err)
		}
	}

	return nil
}

func (s *RedisSubscriber) getEventType(channel string) (int, error) {
	parts := strings.Split(channel, s.prefix)
	if len(parts) < 1 {
		return 0, InvalidChannelError
	}

	eventType, err := strconv.ParseInt(parts[1], 10, 8)
	if err != nil {
		return 0, err
	}

	return int(eventType), nil
}

func (s *RedisSubscriber) errorHandle(err error) {
	baseError := baseerror.NewUndefinedError(err)
	if baseError.HttpStatusCode >= http.StatusInternalServerError {
		s.log.Error(baseError.String())
	}
}

func NewRedisSubscriber(
	log logger.Logger,
	redisClient *redis.Client,
	prefix string,
	eventTypes []int,
	eventHandler EventHandler,
) *RedisSubscriber {
	return &RedisSubscriber{
		log:          log,
		redisClient:  redisClient,
		prefix:       prefix,
		eventTypes:   eventTypes,
		eventHandler: eventHandler,
	}
}
