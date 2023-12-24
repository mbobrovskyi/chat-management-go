package aggregate_root

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/entity"
)

type AggregateRoot[T any] interface {
	entity.Entity[T]
}

type aggregateRoot[T any] struct {
	entity.Entity[T]
}

func Empty[T any]() AggregateRoot[T] {
	return &aggregateRoot[T]{entity.Empty[T]()}
}

func New[T any](id uint64) AggregateRoot[T] {
	return &aggregateRoot[T]{entity.New[T](id)}
}
