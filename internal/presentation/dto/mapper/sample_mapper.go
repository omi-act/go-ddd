package mapper

import (
	"go-ddd/internal/application/command"
	"go-ddd/internal/application/common"
	"go-ddd/internal/presentation/dto/request"
	"go-ddd/internal/presentation/dto/response"
)

// RequestMapper はリクエストDTOとコマンドDTOの変換を行います。
type RequestMapper struct{}

// ToGreetByIdCommand はGreetByIdRequestをGreetByIdCommandに変換します。
// この関数は、プレゼンテーション層のリクエストDTOをアプリケーション層のコマンドDTOに変換します。
func ToGreetByIdCommand(req *request.GreetByIdRequest) (*command.GreetByIdCommand, error) {
	if req == nil {
		return nil, nil
	}

	return command.NewGreetByIdCommand(req.ID), nil
}

// ToGreetByIdResponse はGreetingResultをGreetByIdResponseに変換します。
// この関数は、サービスの結果をプレゼンテーション層のレスポンスDTOに変換します。
func ToGreetByIdResponse(result *common.GreetingResult) *response.GreetByIdResponse {
	return &response.GreetByIdResponse{
		Message: result.Message,
	}
}