package subscriber

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/baseerror"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/logger"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub/event"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/startable"
	"net/http"
	"slices"
)

var _ startable.Startable = (*MemorySubscriber)(nil)

type MemorySubscriber struct {
	log          logger.Logger
	ch           chan event.Event
	eventTypes   []int
	eventHandler EventHandler
	isStarted    bool
}

func (s *MemorySubscriber) Start(ctx context.Context) error {
	if s.isStarted {
		return AlreadyStartedError
	}

	s.isStarted = true
	defer func() {
		s.isStarted = false
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-s.ch:
			if slices.Contains(s.eventTypes, event.Type) {
				if err := s.eventHandler.Handle(event.Type, event.Data); err != nil {
					s.errorHandle(err)
				}
			}
		}
	}
}

func (s *MemorySubscriber) errorHandle(err error) {
	baseError := baseerror.NewUndefinedError(err)
	if baseError.HttpStatusCode >= http.StatusInternalServerError {
		s.log.Error(baseError.String())
	}
}

func NewMemorySubscriber(
	log logger.Logger,
	ch chan event.Event,
	eventTypes []int,
	eventHandler EventHandler,
) *MemorySubscriber {
	return &MemorySubscriber{
		log:          log,
		ch:           ch,
		eventTypes:   eventTypes,
		eventHandler: eventHandler,
	}
}
