package main

import (
	"log"

	"go-ddd/internal/app"
)

// main は Go-DDD アプリケーションのエントリーポイントです。
func main() {
	// サーバー起動
	port := ":9000"
	e := app.Initialize()
	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
