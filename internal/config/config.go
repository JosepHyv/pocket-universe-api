package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabasePort int `env:"DATABASE_PORT"`
	DatabaseHost string `env:"DATABASE_HOST"`
	StoragePort int `env:"STORAGE_PORT"`
	StorageHost string `env:"STORAGE_HOST"`
}

var cfg Config

func LoadConfig() (Config, error){
	err := env.Parse(&cfg)
	if err != nil  {
		return Config{}, fmt.Errorf("Loading environment")
	}
	return cfg, nil
}
