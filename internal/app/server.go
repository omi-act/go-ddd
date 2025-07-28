package app

import (
	"log"
	"os"

	"go-ddd/config"
	"go-ddd/internal/application/services"
	"go-ddd/internal/infrastructure/postgres/repositories"
	"go-ddd/internal/presentation/controllers"

	"github.com/labstack/echo/v4"
)

// Initialize はアプリケーションの初期化を行います。
func Initialize() *echo.Echo {
	configPath := "config/database_prod.yml"
	if IsTestEnvironment() {
		configPath = "../config/database_test.yml"
	}

	return initByConfig(configPath)
}

// initByConfig は指定された設定ファイルパスでアプリケーションの初期化を行います。
func initByConfig(configPath string) *echo.Echo {
	// データベース接続設定を読み込み
	dbConfig, err := config.LoadYAML(configPath)
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
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

// IsTestEnvironment はテスト環境かどうかを判定します。
func IsTestEnvironment() bool {
	// go testコマンドで実行されているかを判定
	for _, arg := range os.Args {
		if arg == "-test.v" || arg == "-test.run" {
			return true
		}
	}

	// TEST_ENVという環境変数でも制御可能
	return os.Getenv("TEST_ENV") == "true"
}
