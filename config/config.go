package config

import (
	"errors"
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
	ErrEmptyServerAddress    = errors.New("empty server address")
	ErrEmptyDBName           = errors.New("empty db name")
	ErrEmptyDBUser           = errors.New("empty db user")
	ErrEmptyDBPass           = errors.New("empty db pass")
	ErrEmptyDBHost           = errors.New("empty db host")
	ErrEmptyDBPort           = errors.New("empty db port")
	ErrEmptyJWTKey           = errors.New("empty JWTKey")
	ErrEmptyDBMigrationsPath = errors.New("empty DBMigrationsPath")
)

func (c *Config) Validate() error {
	if c.ServerAddress == "" {
		return ErrEmptyServerAddress
	}
	if c.DBName == "" {
		return ErrEmptyDBName
	}
	if c.DBUser == "" {
		return ErrEmptyDBUser
	}
	if c.DBPass == "" {
		return ErrEmptyDBPass
	}
	if c.DBHost == "" {
		return ErrEmptyDBHost
	}
	if c.DBPort == "" {
		return ErrEmptyDBPort
	}
	if c.JWTKey == "" {
		return ErrEmptyJWTKey
	}
	if c.DBMigrationsPath == "" {
		return ErrEmptyDBMigrationsPath
	}
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
	return &cfg, nil
}
