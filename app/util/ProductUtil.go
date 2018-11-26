package util

import (
	"../model"
)

var Product1 = model.Product{
	Name: "product1",
	Quantity: 1,
	Sku: "sku1",
	Description: "description1",
	MaximumBorrowedTime: 1,
	MinimumBorrowedTime: 1,
	ProductStatus: model.OPENED,
	PricePerItemPerDay: 1,
	ImageName: "image1",
	CategoryID: 1,
	TenantID: 1,
}

var Product2 = model.Product{
	Name: "product2",
	Quantity: 2,
	Sku: "sku2",
	Description: "description2",
	MaximumBorrowedTime: 1,
	MinimumBorrowedTime: 1,
	ProductStatus: model.OPENED,
	PricePerItemPerDay: 1,
	ImageName: "image1",
	CategoryID: 1,
	TenantID: 1,
}

var Product3_Closed = model.Product{
	Name: "product3",
	Quantity: 3,
	Sku: "sku3",
	Description: "description3",
	MaximumBorrowedTime: 1,
	MinimumBorrowedTime: 1,
	ProductStatus: model.CLOSED,
	PricePerItemPerDay: 1,
	ImageName: "image3",
	CategoryID: 1,
	TenantID: 1,
}

var Products = []model.Product{
	Product1,
	Product2,
}