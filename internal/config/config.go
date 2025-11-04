package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"go-microservice"`
	Port    string `env:"PORT" envDefault:"8080"`

	DBHost     string `env:"DB_HOST" envDefault:"mysql"`
	DBPort     string `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"root"`
	DBName     string `env:"DB_NAME" envDefault:"users_db"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
