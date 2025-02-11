package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Service Service
	Client  Client
}

type Service struct {
	Port string `env:"GATEWAY_SERVICE_PORT"`
}
type Client struct {
	Port string `env:"GATEWAY_CLIENT_PORT"`
}

func MustLoad() *Config {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("can not read env variables: %s", err)
	}
	return cfg
}
