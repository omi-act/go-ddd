package services

import (
	"errors"
	"testing"

	"go-ddd/internal/application/commands"
	"go-ddd/internal/domain/entities"
	"go-ddd/internal/domain/value_objects"
	mocks "go-ddd/test/mocks/infrastructure"

	"github.com/stretchr/testify/assert"
)

// テストデータ定義
type testUserData struct {
	ID    string
	Name  string
	Email string
	Age   int
}

type testCaseData struct {
	Description string
	User        testUserData
	InputID     string
	ExpectedErr string
	ShouldError bool
}

var testCases = struct {
	Valid     testCaseData
	InvalidID testCaseData
	NotFound  testCaseData
}{
	Valid: testCaseData{
		Description: "正常ケース: 有効なユーザーID",
		User: testUserData{
			ID:    "123",
			Name:  "テストユーザー",
			Email: "test@example.com",
			Age:   25,
		},
		InputID:     "123",
		ExpectedErr: "",
		ShouldError: false,
	},
	InvalidID: testCaseData{
		Description: "異常ケース: 無効なユーザーID",
		User:        testUserData{}, // 空のユーザーデータ
		InputID:     "",
		ExpectedErr: "invalid user ID",
		ShouldError: true,
	},
	NotFound: testCaseData{
		Description: "異常ケース: ユーザーが見つからない",
		User:        testUserData{}, // 空のユーザーデータ
		InputID:     "999",
		ExpectedErr: "user not found",
		ShouldError: true,
	},
}

// TestSayHello は GreetingService の SayHello メソッドのテストです。
// このメソッドは "Hello, World!" を返すことを確認します。
func TestSayHello(t *testing.T) {

	service := &GreetingService{}
	got := service.SayHello()

	// testify/assert パッケージを使用したアサーション
	assert.Equal(t, "Hello, World!", got)
}

// TestSayHelloById は GreetingService の SayHelloById メソッドのテストです。
func TestSayHelloById(t *testing.T) {

	// 各テストケースを実行
	// 正常ケース: 有効なユーザーID
	t.Run(testCases.Valid.Description, func(t *testing.T) {
		
		// Arrange
		testdata := testCases.Valid
		mockRepo := mocks.NewMockUserRepository()
		service := NewGreetingService(mockRepo)

		userID, _ := value_objects.NewUserIDFromString(testdata.User.ID)
		user := entities.NewUser(userID.IDNumber, testdata.User.Name, testdata.User.Email, testdata.User.Age)

		// モックの期待値を設定
		mockRepo.On("FindByID", userID).Return(user, nil)

		// Act
		result, err := service.SayHelloById(
			&commands.GreetByIdCommand{
				UserID: testdata.InputID,
			})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result.Message, "Hello, "+testdata.User.Name+"さん(info:")
		assert.Contains(t, result.Message, "ID: "+testdata.User.ID)
		assert.Contains(t, result.Message, "Name: "+testdata.User.Name)
		assert.Contains(t, result.Message, "Email: "+testdata.User.Email)
		assert.Contains(t, result.Message, "Age: 25")

		// モックの期待値が呼ばれたことを確認
		mockRepo.AssertExpectations(t)
	})

	t.Run(testCases.InvalidID.Description, func(t *testing.T) {
		
		// Arrange
		testdata := testCases.InvalidID
		mockRepo := mocks.NewMockUserRepository()
		service := NewGreetingService(mockRepo)

		// Act
		result, err := service.SayHelloById(&commands.GreetByIdCommand{
			UserID: testdata.InputID,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, testdata.ExpectedErr, err.Error())
	})

	t.Run(testCases.NotFound.Description, func(t *testing.T) {
		
		// Arrange
		testdata := testCases.NotFound
		mockRepo := mocks.NewMockUserRepository()
		service := NewGreetingService(mockRepo)
		userID, _ := value_objects.NewUserIDFromString(testdata.InputID)

		// モックの期待値を設定: ユーザーが見つからない場合
		mockRepo.On("FindByID", userID).Return(nil, errors.New(testdata.ExpectedErr))

		// Act
		result, err := service.SayHelloById(&commands.GreetByIdCommand{
			UserID: testdata.InputID, // 存在しないID
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, testdata.ExpectedErr, err.Error())

		// モックの期待値が呼ばれたことを確認
		mockRepo.AssertExpectations(t)
	})
}
