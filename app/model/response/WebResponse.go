package response

type WebResponse struct {
	HttpCode uint
	ErrorCode uint
	Data interface{}
}

func OK(object interface{}) WebResponse {
	return WebResponse {
		HttpCode: 200,
		ErrorCode: 0,
		Data: object,
	}
}

func ERROR(errorCode uint) WebResponse {
	return WebResponse{
		HttpCode: 200,
		ErrorCode: errorCode,
		Data: nil,
	}
}