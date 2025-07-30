package main

import (
	"fmt"
	"os"

	"go-ddd/bootstrap"

	"github.com/joho/godotenv"
)

const (
	envFile = ".env.local"

	// 終了コード定数
	ExitError   = 1 // エラー終了
)

// main は Go-DDD アプリケーションのエントリーポイントです。
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(ExitError)
	}
}

// run は実際のアプリケーションロジックを実行します。
func run() error {
	// .envファイルを読み込み
	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	// サーバー起動
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return fmt.Errorf("SERVER_PORT environment variable is not set")
	}

	e, err := bootstrap.Initialize(envFile)
	if err != nil {
		return fmt.Errorf("failed to initialize application: %w", err)
	}

	if err := e.Start(":" + port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
