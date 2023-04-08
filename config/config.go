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
}

func ParseYAML() (Config, error) {
	body, err := os.ReadFile("./config/config.yml")
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	if err := yaml.Unmarshal(body, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
