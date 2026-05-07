package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Storage Storage
}

type Storage struct {
	Host string
	Port int
}

func Load() (*AppConfig, error) {
	env := strings.TrimSpace(os.Getenv("ENV"))
	if env == "" || env == "development" {
		if err := godotenv.Load(); err != nil && os.IsNotExist(err) {
			return nil, fmt.Errorf("loading development environment file: %w", err)
		}
	}

	env = strings.TrimSpace(os.Getenv("ENV"))
	if env == "" {
		return nil, fmt.Errorf("ENV variable is not set")
	}

	qdrantHost, err := requiredEnv("QDRANT_HOST")
	if err != nil {
		return nil, fmt.Errorf("loading QDRANT_HOST: %w", err)
	}

	qdrantPortStr, err := requiredEnv("QDRANT_PORT")
	if err != nil {
		return nil, fmt.Errorf("loading QDRANT_PORT: %w", err)
	}
	qdrantPort, err := strconv.Atoi(qdrantPortStr)
	if err != nil {
		return nil, fmt.Errorf("parsing QDRANT_PORT: %w", err)
	}

	return &AppConfig{
		Storage: Storage{
			Host: qdrantHost,
			Port: qdrantPort,
		},
	}, nil
}

func requiredEnv(key string) (string, error) {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return "", fmt.Errorf("%s environment variable is required", key)
	}
	return value, nil
}
