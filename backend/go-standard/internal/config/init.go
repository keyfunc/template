package config

import (
	"fmt"
	"os"
	"strings"
)

func Init() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Name: os.Getenv("SERVICE_NAME"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Logger: LogConfig{
			Level:  os.Getenv("LOGGER_LEVEL"),
			Format: os.Getenv("LOGGER_FORMAT"),
		},
		DB: DBConfig{
			Host:    os.Getenv("DB_HOST"),
			Port:    os.Getenv("DB_PORT"),
			User:    os.Getenv("DB_USER"),
			Pwd:     os.Getenv("DB_PASSWORD"),
			Name:    os.Getenv("DB_NAME"),
			SSLMode: os.Getenv("DB_SSLMODE"),
		},
		Redis: RedisConfig{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       os.Getenv("REDIS_DB"),
		},
	}

	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validate(cfg *Config) error {
	var missing []string

	appendMissing := func(name string, value string) {
		if strings.TrimSpace(value) == "" {
			missing = append(missing, name)
		}
	}

	appendMissing("SERVICE_NAME", cfg.Server.Name)
	appendMissing("SERVER_PORT", cfg.Server.Port)

	appendMissing("LOGGER_LEVEL", cfg.Logger.Level)
	appendMissing("LOGGER_FORMAT", cfg.Logger.Format)

	appendMissing("DB_HOST", cfg.DB.Host)
	appendMissing("DB_PORT", cfg.DB.Port)
	appendMissing("DB_USER", cfg.DB.User)
	appendMissing("DB_PASSWORD", cfg.DB.Pwd)
	appendMissing("DB_NAME", cfg.DB.Name)
	appendMissing("DB_SSLMODE", cfg.DB.SSLMode)

	appendMissing("REDIS_ADDR", cfg.Redis.Addr)
	appendMissing("REDIS_PASSWORD", cfg.Redis.Password)
	appendMissing("REDIS_DB", cfg.Redis.DB)

	if len(missing) > 0 {
		return fmt.Errorf("配置不能为空: %s", strings.Join(missing, ", "))
	}

	return nil
}
