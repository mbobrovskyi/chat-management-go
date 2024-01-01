package config

type PubSubType string

const (
	MemoryPubSub PubSubType = "memory"
	RedisPubSub  PubSubType = "redis"
)

func (e PubSubType) String() string {
	return string(e)
}
