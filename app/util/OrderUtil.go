package util

import (
	"../model"
	"time"
)

var Order1 = model.Order{
	BorrowerID: User1.ID,
	ProductID: Product1.ID,
	Quantity: Product1.Quantity,
	DeliveryType: model.PICK_BY_BORROWER,
	RentDurationInDay: 1,
	ReturnTime: time.Now(),
	OrderStatus: model.ON_PROCESS,
	TotalPrice: 15000,
}
var Order2_StatusNotOnProcess = model.Order{
	BorrowerID: User1.ID,
	ProductID: Product1.ID,
	Quantity: Product1.Quantity,
	DeliveryType: model.PICK_BY_BORROWER,
	RentDurationInDay: 1,
	ReturnTime: time.Now(),
	OrderStatus: model.RETRIEVED,
	TotalPrice: 15000,
}