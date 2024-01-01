package config

import "github.com/mbobrovskyi/chat-management-go/pkg/baseconfig"

type Config struct {
	baseconfig.BaseConfig

	DBType     DBType     `env:"DB_TYPE" envDefault:"memory"`
	PubSubType PubSubType `env:"PUB_SUB_TYPE" envDefault:"memory"`

	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
	Redis    RedisConfig    `envPrefix:"REDIS_"`

	GetCurrentUserUrl string `env:"GET_CURRENT_USER_URL" envDefault:"http://localhost:8080/users/current"`
	FindUsersByIdsUrl string `env:"FIND_USERS_BY_IDS_URL" envDefault:"http://localhost:8080/users"`

	ChatPubSubPrefix string `env:"CHAT_PUB_SUB_PREFIX" envDefault:"chat_"`
}

type PostgresConfig struct {
	PostgresUri string `env:"URI"`
}

type RedisConfig struct {
	Addr     string `env:"ADDR" envDefault:"localhost:6379"`
	Password string `env:"PASSWORD"`
	Db       int    `env:"DB"`
}
