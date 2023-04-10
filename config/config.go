package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerAddress    string `yaml:"SERVER_ADDRESS"`
	DBName           string `yaml:"DB_NAME"`
	DBUser           string `yaml:"DB_USER"`
	DBPass           string `yaml:"DB_PASS"`
	DBHost           string `yaml:"DB_HOST"`
	DBPort           string `yaml:"DB_PORT"`
	JWTKey           string `yaml:"JWT_KEY"`
	DBMigrationsPath string `yaml:"DB_MIGRATIONS_PATH"`
}

var (
	ErrEmptyServerAddress = errors.New("empty server address")
)

func (c *Config) Validate() error {
	if c.ServerAddress == "" {
		return ErrEmptyServerAddress
	}
	// if ser
	return nil
}

func ParseYAML() (*Config, error) {
	body, err := os.ReadFile("./config/.config.yml")
	if err != nil {
		return nil, err
	}
	cfg := Config{}
	if err := yaml.Unmarshal(body, &cfg); err != nil {
		return nil, err
	}
	fmt.Println(cfg.DBHost)
	return &cfg, nil
}
