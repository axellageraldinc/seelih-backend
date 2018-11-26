package util

import (
	"../model/request"
)

var ConfirmProductCancellationRequest1 = request.ConfirmProductCancellationRequest{
	OrderId: Order1.ID,
}

var ConfirmProductCancellationRequest2_OrderIdNotFound = request.ConfirmProductCancellationRequest{
	OrderId: 999,
}