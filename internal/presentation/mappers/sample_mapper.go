package mappers

import (
	"go-ddd/internal/application/commands"
	"go-ddd/internal/application/common"
	"go-ddd/internal/presentation/requests"
	"go-ddd/internal/presentation/responses"
)

// RequestMapper はリクエストDTOとコマンドDTOの変換を行います。
type RequestMapper struct{}

// ToGreetByIdCommand はGreetByIdRequestをGreetByIdCommandに変換します。
// この関数は、プレゼンテーション層のリクエストDTOをアプリケーション層のコマンドDTOに変換します。
func ToGreetByIdCommand(req *requests.GreetByIdRequest) (*commands.GreetByIdCommand, error) {
	if req == nil {
		return nil, nil
	}

	return commands.NewGreetByIdCommand(req.ID), nil
}

// ToGreetByIdResponse はGreetingResultをGreetByIdResponseに変換します。
// この関数は、サービスの結果をプレゼンテーション層のレスポンスDTOに変換します。
func ToGreetByIdResponse(result *common.GreetingResult) *responses.GreetByIdResponse {
	return &responses.GreetByIdResponse{
		Message: result.Message,
	}
}