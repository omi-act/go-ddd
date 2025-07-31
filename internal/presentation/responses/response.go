package responses

// GreetByIdResponse はGreetByIdのAPIレスポンスを表すDTOです。
type GreetByIdResponse struct {
	Message string `json:"message"`
}

// NewGreetByIdResponse は新しいGreetByIdResponseを作成します。
func NewGreetByIdResponse(message string, status string, code int) *GreetByIdResponse {
	return &GreetByIdResponse{
		Message: message,
	}
}
