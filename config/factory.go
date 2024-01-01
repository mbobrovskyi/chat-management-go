package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	c := &Config{}

	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}
