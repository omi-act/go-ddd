package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestSayHello は GreetingService の SayHello メソッドのテストです。
// このメソッドは "Hello, World!" を返すことを確認します。
func TestSayHello(t *testing.T) {

	service := &GreetingService{}
	got := service.SayHello()
	// want := "Hello, World!"
	// want := "FAILED"

	// 標準テストライブラリ版
	// if got != want {
	// 	t.Errorf("SayHello() = %v, want %v", got, want)
	// } else {
	// 	t.Logf("SayHello() = %v, as expected", got)
	// 	println("Test passed successfully!")
	// }

	// testify/assert パッケージを使用したアサーション
	assert.Equal(t, "Hello, World!", got)
	//assert.Equal(t, "FAILED", got)
}
