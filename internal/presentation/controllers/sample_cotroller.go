package controllers

import (
	"go-ddd/internal/application/interfaces"
	"go-ddd/internal/presentation/mappers"
	"go-ddd/internal/presentation/requests"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SampleController はサンプルのコントローラです。
type SampleController struct {
	service       interfaces.GreetingService
}

// NewSampleController は SampleController の新しいインスタンスを作成します。
func NewSampleController(e *echo.Echo, service interfaces.GreetingService) *SampleController {
	controller := &SampleController{
		service:       service,
	}

	// ルーティングの設定
	e.GET("/sample/greet", controller.Greet)
	e.GET("/sample/greetById/:id", controller.GreetById)
	e.Use(middleware.Recover())

	return controller
}

// Greet はサンプルの挨拶を行うハンドラーです。
func (sc *SampleController) Greet(c echo.Context) error {
	response := sc.service.SayHello()
	return c.JSON(http.StatusOK, response)
}

// GreetById はIDに基づいてサンプルの挨拶を行うハンドラーです。
func (sc *SampleController) GreetById(c echo.Context) error {

	// パラメータの取得
	id := c.Param("id")
	
	//リクエストDTOをコマンドDTOに変換
	req := requests.NewGreetByIdRequest(id)
	cmd, err := mappers.ToGreetByIdCommand(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// サービスの呼び出し
	result, err := sc.service.SayHelloById(cmd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// レスポンスの変換
	response := mappers.ToGreetByIdResponse(result)
	return c.JSON(http.StatusOK, response)
}
