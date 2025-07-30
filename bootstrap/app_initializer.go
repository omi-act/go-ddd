package bootstrap

import (
	"fmt"
	"os"
	"strconv"

	"go-ddd/internal/application/services"
	"go-ddd/internal/infrastructure/postgres"
	"go-ddd/internal/infrastructure/postgres/repositories"
	"go-ddd/internal/presentation/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// Initialize はアプリケーションの初期化を行います。
func Initialize(envFile string) (*echo.Echo, error) {

	// .envファイルを読み込み
	if err := godotenv.Load(envFile); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// データベース接続設定を環境変数から読み込み
	dbConfig, err := loadDBEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config from environment variables: %w", err)
	}

	// データベース接続
	gormDb, err := dbConfig.CreateDbConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// リポジトリとサービスの初期化
	userRepository := repositories.NewGormUserRepository(gormDb)
	greetingService := services.NewGreetingService(userRepository)

	e := echo.New()
	controllers.NewSampleController(e, greetingService)

	return e, nil
}

// loadDBEnv は環境変数からデータベース接続設定を読み込みます。
func loadDBEnv() (*postgres.DatabaseConfig, error) {

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

	return &postgres.DatabaseConfig{
		Host:     envVals["DB_HOST"],
		Port:     port,
		User:     envVals["DB_USER"],
		Password: envVals["DB_PASSWORD"],
		DBName:   envVals["DB_NAME"],
		SSLMode:  envVals["DB_SSLMODE"],
	}, nil
}
