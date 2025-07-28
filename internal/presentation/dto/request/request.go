package request

// GreetByIdRequest はIDに基づく挨拶リクエストのDTOです。
type GreetByIdRequest struct {
	ID string
}

// NewGreetByIdRequest は新しいGreetByIdRequestを作成します。
func NewGreetByIdRequest(id string) *GreetByIdRequest {
	return &GreetByIdRequest{
		ID: id,
	}
}
