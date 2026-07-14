package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabasePort int `env:"DATABASE_PORT"`
	DatabaseHost string `env:"DATABASE_HOST"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	DatabaseUser string `env:"DATABASE_USER"`
	DatabaseCollection string `env:"DATABASE_COLLECTION"`
	DatabaseName string `env:"DATABASE_NAME"`
	StoragePort int `env:"STORAGE_PORT"`
	StorageHost string `env:"STORAGE_HOST"`
	StorageUser string `env:"STORAGE_USER"`
	StoragePassword string `env:"STORAGE_PASSWORD"`
	StorageBucket string `env:"STORAGE_BUCKET"`
	AppPort int `env:"APP_PORT"`
}


func LoadConfig() (Config, error){
	cfg := Config{}
	err := godotenv.Load()
	if err != nil {
		return cfg, fmt.Errorf("Error: No se encontró archivo .env, leyendo variables del sistema directamente")
	}
	err = env.Parse(&cfg)
	if err != nil  {
		return Config{}, fmt.Errorf("Loading environment")
	}
	return cfg, nil
}
