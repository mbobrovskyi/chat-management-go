package startable

import "context"

type Startable interface {
	Start(ctx context.Context) error
}
