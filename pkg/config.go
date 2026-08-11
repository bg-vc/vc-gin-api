package pkg

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Address   string
	RunMode   string
	JWTSecret string
	JWTExpiry time.Duration
	Database  DBConfig
	Redis     RedisConfig
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("load .env: %w", err)
	}

	cfg := Config{
		Address:   env("VC_HTTP_ADDRESS", ":8080"),
		RunMode:   env("VC_GIN_MODE", "debug"),
		JWTSecret: os.Getenv("VC_JWT_SECRET"),
		Database: DBConfig{
			Host:     env("VC_DB_HOST", "127.0.0.1"),
			Port:     env("VC_DB_PORT", "3306"),
			Name:     env("VC_DB_NAME", "demo"),
			User:     env("VC_DB_USER", "vcgin"),
			Password: os.Getenv("VC_DB_PASSWORD"),
		},
		Redis: RedisConfig{
			Enabled:  envBool("VC_REDIS_ENABLED", false),
			Address:  env("VC_REDIS_ADDRESS", "127.0.0.1:6379"),
			Password: os.Getenv("VC_REDIS_PASSWORD"),
			DB:       envInt("VC_REDIS_DB", 0),
		},
	}

	var err error
	cfg.JWTExpiry, err = time.ParseDuration(env("VC_JWT_EXPIRY", "30m"))
	if err != nil {
		return Config{}, fmt.Errorf("parse VC_JWT_EXPIRY: %w", err)
	}
	if len(cfg.JWTSecret) < 32 {
		return Config{}, fmt.Errorf("VC_JWT_SECRET must contain at least 32 characters")
	}
	if cfg.Database.Password == "" {
		return Config{}, fmt.Errorf("VC_DB_PASSWORD is required")
	}

	return cfg, nil
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func envBool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func envInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
