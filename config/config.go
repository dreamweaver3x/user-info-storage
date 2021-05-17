package config

import (
	"fmt"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	AppName string `envconfig:"APP_NAME"`
	DSN     string `envconfig:"DB_DSN"`
	Port    uint   `envconfig:"PORT,default=8080"`
}

func (c *Config) ListenAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

func Load() (*Config, error) {
	var c *Config
	if err := envconfig.Init(&c); err != nil {
		return nil, err
	}
	return c, nil
}
