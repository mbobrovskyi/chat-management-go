package configs

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment    Environment `env:"ENVIRONMENT" envDefault:"development"`
	HttpServerAddr string      `env:"HTTP_SERVER_ADDR" envDefault:"0.0.0.0:8080"`
	LogLevel       string      `env:"LOG_LEVEL" envDefault:"debug"`

	PostgresUri string `env:"POSTGRES_URI" envDefault:"postgresql://postgres:postgres@localhost:5432/chat?sslmode=disable"`

	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDb       int    `env:"REDIS_DB"`

	ChatPubSubPrefix string `env:"CHAT_PUB_SUB_PREFIX" envDefault:"chat_"`
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}
