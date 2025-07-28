package common

// GreetingResult は挨拶処理の結果を表すDTOです。
type GreetingResult struct {
	Message string
}

// NewGreetingResult は新しいGreetingResultを作成します。
func NewGreetingResult(message string, success bool) *GreetingResult {
	return &GreetingResult{
		Message: message,
	}
}