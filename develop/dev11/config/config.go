package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

type (
	Config struct {
		PG   `yaml:"postgres"`
		HTTP `yaml:"http"`
	}
	PG struct {
		ConnStr     string `env-required:"true" env:"PG_CONN_STR"`
		MaxPoolSize int    `yaml:"max_pool_size"`
	}
	Log struct {
		Level string `yaml:"level"`
	}
	HTTP struct {
		Address string `yaml:"address"`
	}
)

func MustLoad(cfgPath string) *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var cfg Config

	if err = cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("Error reading config .env file")
	}

	return &cfg
}
