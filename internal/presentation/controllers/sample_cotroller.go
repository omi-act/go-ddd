package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"go-ddd/internal/application/interfaces"
)

// SampleController はサンプルのコントローラです。
type SampleController struct {
	service interfaces.GreetingService
}

// NewSampleController は SampleController の新しいインスタンスを作成します。
func NewSampleController(e *echo.Echo, service interfaces.GreetingService) *SampleController {
	controller := &SampleController{
		service: service,
	}

	// ルーティングの設定
	e.GET("/sample/greet", controller.Greet)
	e.Use(middleware.Recover())

	return controller
}

// Greet はサンプルの挨拶を行うハンドラーです。
func (sc *SampleController) Greet(c echo.Context) error  {
	response := sc.service.SayHello()
	return c.JSON(http.StatusOK, response)
}