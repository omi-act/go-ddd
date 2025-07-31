package bootstrap

import (
	"fmt"
	"os"
	"strconv"

	"go-ddd/internal/application/services"
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
	gormDb, err := dbConfig.OpenConnection()
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
func loadDBEnv() (*repositories.DatabaseConfig, error) {

	// 環境変数読み込み
	requiredEnvs := map[string]string{
		"DB_HOST":     "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
		"DB_PORT":     "",
	}
	for envKey := range requiredEnvs {
		val := os.Getenv(envKey)
		if val == "" {
			return nil, fmt.Errorf("%s environment variable is required", envKey)
		}

		requiredEnvs[envKey] = val
	}
	port, err := strconv.Atoi(requiredEnvs["DB_PORT"])
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %w", err)
	}

	return &repositories.DatabaseConfig{
		Host:     requiredEnvs["DB_HOST"],
		Port:     port,
		User:     requiredEnvs["DB_USER"],
		Password: requiredEnvs["DB_PASSWORD"],
		DBName:   requiredEnvs["DB_NAME"],
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}, nil
}
