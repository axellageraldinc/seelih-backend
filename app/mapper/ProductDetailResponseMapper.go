package mapper

import (
	"../model"
	"../model/response"
)

func ToProductDetailResponse(product model.Product) response.ProductDetailResponse {
	return response.ProductDetailResponse{
		Id: product.ID,
		TenantId: product.TenantID,
		CategoryId: product.CategoryID,
		Sku: product.Sku,
		Name: product.Name,
		Quantity: product.Quantity,
		PricePerItemPerDay: product.PricePerItemPerDay,
		Description: product.Description,
		UploadedTime: product.CreatedAt,
		MinimumBorrowedTime: product.MinimumBorrowedTime,
		MaximumBorrowedTime: product.MaximumBorrowedTime,
		ImageUrl: model.IMAGE_URL_PREFIX + product.ImageName,
	}
}