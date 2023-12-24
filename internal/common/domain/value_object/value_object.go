package value_object

type ValueObject[T any] interface {
	Equals(other T) bool
}
