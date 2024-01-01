package publisher

import "context"

type Publisher interface {
	Publish(ctx context.Context, eventType int, data any) error
}
