package resqueExporter

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	GuardIntervalMillis int64        `env:"GUARD_INTERVAL_MILLIS"`
	ResqueNamespace     string       `env:"RESQUE_NAMESPACE"`
	Redis               *RedisConfig
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int64  `env:"REDIS_DB"`
}

func loadConfigFromEnv() (*Config, error) {
	var config Config

	// Load GuardIntervalMillis
	guardIntervalMillisStr := os.Getenv("GUARD_INTERVAL_MILLIS")
	if guardIntervalMillisStr != "" {
		guardIntervalMillis, err := strconv.ParseInt(guardIntervalMillisStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse GUARD_INTERVAL_MILLIS; err: %s", err)
		}
		config.GuardIntervalMillis = guardIntervalMillis
	}

	// Load ResqueNamespace
	config.ResqueNamespace = os.Getenv("RESQUE_NAMESPACE")

	// Load RedisConfig
	config.Redis = &RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	portStr := os.Getenv("REDIS_PORT")
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse REDIS_PORT; err: %s", err)
		}
		config.Redis.Port = port
	}

	dbStr := os.Getenv("REDIS_DB")
	if dbStr != "" {
		db, err := strconv.ParseInt(dbStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse REDIS_DB; err: %s", err)
		}
		config.Redis.DB = db
	}

	return &config, nil
}