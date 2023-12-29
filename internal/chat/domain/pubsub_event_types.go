package domain

const (
	CreateMessagePubSubEventType uint8 = 1
)

func GetAllPubSubEventTypes() []uint8 {
	return []uint8{CreateMessagePubSubEventType}
}
