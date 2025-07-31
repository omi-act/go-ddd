package mocks

import (
	"go-ddd/internal/domain/entities"
	"go-ddd/internal/domain/value_objects"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository はtestify/mockを使用したテスト用のモックリポジトリです
type MockUserRepository struct {
	mock.Mock
}

// NewMockUserRepository は新しいMockUserRepositoryインスタンスを作成します
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

// FindByID はユーザーIDでユーザーを検索します
func (m *MockUserRepository) FindByID(userID *value_objects.UserID) (*entities.User, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

// FindAll は全てのユーザーを取得します
func (m *MockUserRepository) FindAll(limit, offset int) ([]*entities.User, error) {
	args := m.Called(limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.User), args.Error(1)
}
