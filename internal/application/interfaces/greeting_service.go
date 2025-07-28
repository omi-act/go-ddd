package interfaces

import (
	"go-ddd/internal/application/command"
	"go-ddd/internal/application/common"
)

// GreetingService は挨拶を提供するサービスのインターフェースです。
type GreetingService interface {
	SayHello() string
	SayHelloById(cmd *command.GreetByIdCommand) (*common.GreetingResult, error)
}
