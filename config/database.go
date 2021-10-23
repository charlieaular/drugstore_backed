package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBUserName string `envconfig:"DB_USERNAME" default:"postgres"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"postgres"`
}

func NewConfig() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
