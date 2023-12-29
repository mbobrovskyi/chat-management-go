package message

type Repository interface {
	GetMessages() ([]Message, uint64, error)
}
