package domain

const (
	CreateMessagePubSubEventType = 1
)

func GetAllPubSubEventTypes() []int {
	return []int{CreateMessagePubSubEventType}
}
