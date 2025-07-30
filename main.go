package main

import (
	"os"
	"fmt"

	"go-ddd/internal/app"

	"github.com/joho/godotenv"
)

const (
	envFile = ".env.local"
)

// main は Go-DDD アプリケーションのエントリーポイントです。
func main() {
	// .envファイルを読み込み
	if err := godotenv.Load(envFile); err != nil {
		panic(fmt.Sprintf("Warning: .env file not found or could not be loaded: %v", err))
	}

	// サーバー起動
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		panic("SERVER_PORT environment variable is not set")
	}
	e, err := app.Initialize()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize application: %v", err))
	}
	if err := e.Start(":" + port); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
