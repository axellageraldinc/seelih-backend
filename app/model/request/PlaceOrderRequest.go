package request

type PlaceOrderRequest struct {
	ProductId uint
	BorrowerId uint
	Quantity uint
	Duration uint
	DeliveryType string
}