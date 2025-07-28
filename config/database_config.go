package config

import (
	"fmt"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v3"

	"go-ddd/internal/infrastructure/postgres"
)

// DatabaseConfig はデータベース設定を表す構造体です
type DatabaseConfig struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

// LoadYAML は設定ファイルからデータベース設定を読み込みます
func LoadYAML(configPath string) (*DatabaseConfig, error) {
	// 絶対パスに変換
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	// ファイルを読み込み
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config DatabaseConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// ToConnection は DatabaseConfig を postgres.Connection に変換します。
func (c *DatabaseConfig) ToConnection() *postgres.Connection {
	return &postgres.Connection{
		Host:     c.Database.Host,
		Port:     c.Database.Port,
		User:     c.Database.User,
		Password: c.Database.Password,
		DBName:   c.Database.DBName,
	}
}