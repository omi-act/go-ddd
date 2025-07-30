package main

import (
	"log"
	"os"

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
		log.Printf("Warning: .env file not found or could not be loaded: %v", err)
	}

	// サーバー起動
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		panic("SERVER_PORT environment variable is not set")
	}

	e := app.Initialize()
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
