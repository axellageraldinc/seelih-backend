package util

import (
	"../model"
	"../model/request"
)

var PlaceOrderRequest1_OrderIdNotFound = request.PlaceOrderRequest{
	DeliveryType: model.PICK_BY_BORROWER,
	Quantity: 1,
	BorrowerId: User2.ID,
	ProductId: 999,
	Duration: 1,
}

var PlaceOrderRequest2_OrderNotAvailableForRenting = request.PlaceOrderRequest{
	DeliveryType: model.PICK_BY_BORROWER,
	Quantity: 1,
	BorrowerId: User2.ID,
	ProductId: 1,
	Duration: 1,
}

var PlaceOrderRequest3_BorrowerIsTenant = request.PlaceOrderRequest{
	DeliveryType: model.PICK_BY_BORROWER,
	Quantity: 1,
	BorrowerId: 1,
	ProductId: 1,
	Duration: 1,
}

var PlaceOrderRequest4_DurationDoesntMeetMinimum= request.PlaceOrderRequest{
	DeliveryType: model.PICK_BY_BORROWER,
	Quantity: 1,
	BorrowerId: 9,
	ProductId: 1,
	Duration: 0,
}

var PlaceOrderRequest5_DurationExceedsMaximum= request.PlaceOrderRequest{
	DeliveryType: model.PICK_BY_BORROWER,
	Quantity: 1,
	BorrowerId: 9,
	ProductId: 1,
	Duration: 100,
}

var PlaceOrderRequest6_RequestedQuantityExceedsCurrentQuantity= request.PlaceOrderRequest{
	DeliveryType: model.PICK_BY_BORROWER,
	Quantity: 100,
	BorrowerId: 9,
	ProductId: 1,
	Duration: 1,
}

var PlaceOrderRequest7_BorrowerIdNotFound= request.PlaceOrderRequest{
	DeliveryType: model.PICK_BY_BORROWER,
	Quantity: 1,
	BorrowerId: 999,
	ProductId: 1,
	Duration: 1,
}

var PlaceOrderRequest8 = request.PlaceOrderRequest{
	DeliveryType: model.PICK_BY_BORROWER,
	Quantity: 1,
	BorrowerId: 9,
	ProductId: Product2.ID,
	Duration: 1,
}