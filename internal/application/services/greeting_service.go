package services

import (
	"errors"
	"go-ddd/internal/application/command"
	"go-ddd/internal/application/common"
	"go-ddd/internal/application/interfaces"
	"go-ddd/internal/domain/repositories"
	"go-ddd/internal/domain/value_objects"
)

// GreetingService は挨拶を提供するサービスの実装です。
type GreetingService struct {
	repository repositories.UserRepository
}

// NewGreetingService は GreetingService の新しいインスタンスを作成します。
func NewGreetingService(rep repositories.UserRepository) interfaces.GreetingService {
	return &GreetingService{
		repository: rep,
	}
}

// SayHello は挨拶を返すメソッドです。
func (s *GreetingService) SayHello() string {
	return "Hello, World!"
}

// SayHelloById はIDに基づいて挨拶を返すメソッドです。
func (s *GreetingService) SayHelloById(cmd *command.GreetByIdCommand) (*common.GreetingResult, error) {

	// コマンドの検証
	userID, err := value_objects.NewUserIDFromString(cmd.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	// ユーザーをリポジトリから取得
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	message := "Hello, " + user.Show() + "!"
	return &common.GreetingResult{
		Message: message,
	}, nil
}
