package util

import (
	"../model/request"
)

var ConfirmProductReturnRequest1 = request.ConfirmProductReturnRequest{
	OrderId: Order1.ID,
}
var ConfirmProductReturnRequest2_OrderIdNotFound = request.ConfirmProductReturnRequest{
	OrderId: 999,
}