package entity

type Entity[T any] interface {
	GetId() uint64
	Equals(other Entity[T]) bool
}

type entity[T any] struct {
	Id uint64 `json:"id"`
}

func (e *entity[T]) GetId() uint64 {
	return e.Id
}

func (e *entity[T]) Equals(other Entity[T]) bool {
	if other == nil {
		return false
	}

	if e.GetId() == 0 && other.GetId() == 0 {
		return false
	}

	return e.GetId() == other.GetId()
}

func Empty[T any]() Entity[T] {
	return &entity[T]{0}
}

func New[T any](id uint64) Entity[T] {
	return &entity[T]{id}
}
