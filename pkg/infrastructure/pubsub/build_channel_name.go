package pubsub

import "fmt"

func BuildChannelName(prefix string, eventType int) string {
	return fmt.Sprintf("%s%d", prefix, eventType)
}
