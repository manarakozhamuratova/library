package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerAddress string `yaml:"SERVER_ADDRESS"`
	DBName        string `yaml:"DB_NAME"`
	DBUser        string `yaml:"DB_USER"`
	DBPass        string `yaml:"DB_PASS"`
	JWTKey        string `envconfig:"JWT_KEY" default:"supersecret"`
}

func ParseYAML() (*Config, error) {
	body, err := os.ReadFile("./config/config.yml")
	if err != nil {
		return nil, err
	}
	cfg := Config{}
	if err := yaml.Unmarshal(body, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
