package mapper

import (
	. "../model"
	. "../model/response"
)

type IProductDetailResponseMapper interface {
	ToProductDetailResponse(product Product) ProductDetailResponse
}

type ProductDetailResponseMapper struct{}

func (productDetailResponseMapper *ProductDetailResponseMapper) ToProductDetailResponse(product Product) ProductDetailResponse {
	return ProductDetailResponse{
		Id:                  product.ID,
		TenantId:            product.TenantID,
		CategoryId:          product.CategoryID,
		Sku:                 product.Sku,
		Name:                product.Name,
		Quantity:            product.Quantity,
		PricePerItemPerDay:  product.PricePerItemPerDay,
		Description:         product.Description,
		UploadedTime:        product.CreatedAt,
		MinimumBorrowedTime: product.MinimumBorrowedTime,
		MaximumBorrowedTime: product.MaximumBorrowedTime,
		ImageUrl:            IMAGE_URL_PREFIX + product.ImageName,
	}
}