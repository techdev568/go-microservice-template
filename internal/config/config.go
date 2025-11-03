package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"go-microservice"`
	Port    string `env:"PORT" envDefault:"8080"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
