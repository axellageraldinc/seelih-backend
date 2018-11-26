package util

import (
	"../model"
	"../model/response"
)

var ProductDetailResponse1 = response.ProductDetailResponse{
	Id: Product1.ID,
	Name: Product1.Name,
	UploadedTime: Product1.CreatedAt,
	ImageUrl: model.IMAGE_URL_PREFIX + Product1.ImageName,
	PricePerItemPerDay: Product1.PricePerItemPerDay,
	CategoryId: Product1.CategoryID,
	TenantId: Product1.TenantID,
	Quantity: Product1.Quantity,
	Sku: Product1.Sku,
	Description: Product1.Description,
	MaximumBorrowedTime: Product1.MaximumBorrowedTime,
	MinimumBorrowedTime: Product1.MinimumBorrowedTime,
}