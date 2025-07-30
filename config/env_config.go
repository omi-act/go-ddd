package config

import (
	"fmt"
	"os"
	"strconv"

	"go-ddd/internal/infrastructure/postgres"
)

// EnviromentType はアプリケーションの実行環境を表す列挙型です。
type EnviromentType int

// EnvConfig は環境変数からの設定を表す構造体です
type EnvConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LoadEnv は環境変数からデータベース設定を読み込みます
func LoadEnv() (*EnvConfig, error) {

	// 環境変数読み込み
	envVals := map[string]string{
		"DB_HOST":     "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
		"DB_PORT":     "",
		"DB_SSLMODE":  "disable", // デフォルト値
	}
	for envKey := range envVals {
		val := os.Getenv(envKey)
		defaultVal := envVals[envKey]
		if val == "" && defaultVal == "" {
			return nil, fmt.Errorf("%s environment variable is required", envKey)
		}

		if val == "" {
			val = defaultVal
		}
		envVals[envKey] = val
	}
	port, err := strconv.Atoi(envVals["DB_PORT"])
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %w", err)
	}

	return &EnvConfig{
		Host:     envVals["DB_HOST"],
		Port:     port,
		User:     envVals["DB_USER"],
		Password: envVals["DB_PASSWORD"],
		DBName:   envVals["DB_NAME"],
		SSLMode:  envVals["DB_SSLMODE"],
	}, nil
}

// ToConnection は EnvConfig を postgres.Connection に変換します
func (c *EnvConfig) ToConnection() *postgres.Connection {
	return &postgres.Connection{
		Host:     c.Host,
		Port:     c.Port,
		User:     c.User,
		Password: c.Password,
		DBName:   c.DBName,
		SSLMode:  c.SSLMode,
	}
}
