package services

import (
	"go-ddd/internal/application/interfaces"
)

// GreetingService は挨拶を提供するサービスの実装です。
type GreetingService struct {
}

// NewGreetingService は GreetingService の新しいインスタンスを作成します。
func NewGreetingService() interfaces.GreetingService {
	return &GreetingService{}
}

// SayHello は挨拶を返すメソッドです。
func (s *GreetingService) SayHello() string {
	return "Hello, World!"
}