package app

import (
	"log"

	"go-ddd/config"
	"go-ddd/internal/application/services"
	"go-ddd/internal/infrastructure/postgres/repositories"
	"go-ddd/internal/presentation/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// Initialize はアプリケーションの初期化を行います。
func Initialize() *echo.Echo {

	// .envファイルを読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or could not be loaded: %v", err)
	}

	// データベース接続設定を環境変数から読み込み
	dbConfig, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load database config from environment variables: %v", err)
	}

	// データベース接続
	gormDb, err := dbConfig.ToConnection().CreateDbConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// リポジトリとサービスの初期化
	userRepository := repositories.NewGormUserRepository(gormDb)
	greetingService := services.NewGreetingService(userRepository)

	e := echo.New()
	controllers.NewSampleController(e, greetingService)

	return e
}
