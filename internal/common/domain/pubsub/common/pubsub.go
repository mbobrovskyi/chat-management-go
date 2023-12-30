package common

import "fmt"

func BuildChannelName(prefix string, eventType uint8) string {
	return fmt.Sprintf("%s%d", prefix, eventType)
}
