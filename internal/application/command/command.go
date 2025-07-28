package command

import (
)

// GreetByIdCommand はIDに基づく挨拶処理のコマンドです。
type GreetByIdCommand struct {
	UserID string
}

// NewGreetByIdCommand は新しいGreetByIdCommandを作成します。
func NewGreetByIdCommand(id string) *GreetByIdCommand {
	return &GreetByIdCommand{
		UserID: id,
	}
}