package config

import (
	"flag"
	"fmt"
	"time"
)

type config struct {
	Host    string
	Port    string
	Timeout string
}

func (c *Config) Validate() error {
	if c.Host == "" || c.Port == "" {
		return fmt.Errorf("validation error, both host and port must be non-empty")
	}
	if _, err := time.ParseDuration(c.Timeout); err != nil {
		return fmt.Errorf("validation error, invalid timeout: %w", err)
	}
	return nil
}

func MustLoad() *Config {
	var cfg Config
	flag.StringVar(&cfg.Host, "host", "localhost", "host for telnet connection")
	flag.StringVar(&cfg.Port, "port", "23", "port for telnet connection")
	flag.StringVar(&cfg.Timeout, "timeout", "3s", "connection timeout")

	flag.Parse()

	return &cfg
}
