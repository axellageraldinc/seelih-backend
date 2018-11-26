package util

import (
	"../model"
	"../model/response"
)

var UploadedProductResponse1 = response.UploadedProductResponse{
	PricePerItemPerDay: Product1.PricePerItemPerDay,
	ImageUrl: model.IMAGE_URL_PREFIX + Product1.ImageName,
	UploadedTime: Product1.CreatedAt,
	Name: Product1.Name,
	Id: Product1.ID,
}
var UploadedProductResponse2 = response.UploadedProductResponse{
	PricePerItemPerDay: Product2.PricePerItemPerDay,
	ImageUrl: model.IMAGE_URL_PREFIX + Product2.ImageName,
	UploadedTime: Product2.CreatedAt,
	Name: Product2.Name,
	Id: Product2.ID,
}
var UploadedProductsResponse = []response.UploadedProductResponse{
	UploadedProductResponse1,
	UploadedProductResponse2,
}