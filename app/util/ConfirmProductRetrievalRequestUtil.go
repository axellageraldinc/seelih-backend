package util

import (
	"../model/request"
)

var ConfirmProductRetrievalRequest1 = request.ConfirmProductRetrievalRequest{
	OrderId: Order1.ID,
}
var ConfirmProductRetrievalRequest2_OrderIdNotFound = request.ConfirmProductRetrievalRequest{
	OrderId: 999,
}