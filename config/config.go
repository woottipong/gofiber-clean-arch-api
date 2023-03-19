package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB     DatabaseConfig
	Server ServerConfig
}

type DatabaseConfig struct {
	Host     string `required:"true" envconfig:"DB_HOST"`
	Port     string `required:"true" envconfig:"DB_PORT"`
	Username string `required:"true" envconfig:"DB_USERNAME"`
	Password string `required:"true" envconfig:"DB_PASSWORD"`
	Name     string `required:"true" envconfig:"DB_NAME"`
}

type ServerConfig struct {
	Host string `default:"" envconfig:"SERVER_HOST"`
	Port string `default:"3000" envconfig:"SERVER_PORT"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
