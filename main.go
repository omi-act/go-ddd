package main

import (
	"github.com/labstack/echo/v4"
	"go-ddd/internal/presentation/controllers"
	"go-ddd/internal/application/services"
	"log"
)

// main は Go-DDD アプリケーションのエントリーポイントです。
func main() {

	// 初期化.サービス
	greetingService := services.NewGreetingService()
	
	// 初期化.コントローラ
	e := echo.New()
	controllers.NewSampleController(e, greetingService)
	
	// サーバー起動
	port := ":9000"
	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
