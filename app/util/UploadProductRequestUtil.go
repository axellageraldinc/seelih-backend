package util

import (
	"../model/request"
)

var UploadProductPricePerItemPerDayIsZeroOrBelow = request.UploadProductRequest{
	Name: "product1",
	PricePerItemPerDay: 0,
	MinimumBorrowedTime: 1,
	MaximumBorrowedTime: 1,
	Description: "description1",
	Sku: "sku1",
	Quantity: 1,
	CategoryId: 1,
	TenantId: 1,
}
var UploadProductPriceQuantityIsZeroOrBelow = request.UploadProductRequest{
	Name: "product1",
	PricePerItemPerDay: 1000,
	MinimumBorrowedTime: 1,
	MaximumBorrowedTime: 1,
	Description: "description1",
	Sku: "sku1",
	Quantity: 0,
	CategoryId: 1,
	TenantId: 1,
}
var UploadProduct1 = request.UploadProductRequest{
	Name: "product1",
	PricePerItemPerDay: 1000,
	MinimumBorrowedTime: 1,
	MaximumBorrowedTime: 1,
	Description: "description1",
	Sku: "sku1",
	Quantity: 1,
	CategoryId: 1,
	TenantId: 1,
}