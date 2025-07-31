package interfaces

import (
	"go-ddd/internal/application/commands"
	"go-ddd/internal/application/common"
)

// GreetingService は挨拶を提供するサービスのインターフェースです。
type GreetingService interface {
	SayHello() string
	SayHelloById(cmd *commands.GreetByIdCommand) (*common.GreetingResult, error)
}
