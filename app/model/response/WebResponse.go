package response

import "net/http"

type WebResponse struct {
	HttpCode uint
	ErrorCode uint
	Data interface{}
}

func OK(object interface{}) WebResponse {
	return WebResponse {
		HttpCode: http.StatusOK,
		ErrorCode: 0,
		Data: object,
	}
}

func ERROR(errorCode uint) WebResponse {
	return WebResponse{
		HttpCode: http.StatusOK,
		ErrorCode: errorCode,
		Data: nil,
	}
}