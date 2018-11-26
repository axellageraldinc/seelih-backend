package util

import (
	"../model"
	"../model/response"
)

var AvailableProductForRentingResponse1 = response.AvailableProductForRentingResponse{
	PricePerItemPerDay: Product1.PricePerItemPerDay,
	Name: Product1.Name,
	Id: Product1.ID,
	ImageUrl: model.IMAGE_URL_PREFIX + Product1.ImageName,
	UploadedTime: Product1.CreatedAt,
}

var AvailableProductForRentingResponse2 = response.AvailableProductForRentingResponse{
	PricePerItemPerDay: Product2.PricePerItemPerDay,
	Name: Product2.Name,
	Id: Product2.ID,
	ImageUrl: model.IMAGE_URL_PREFIX + Product2.ImageName,
	UploadedTime: Product2.CreatedAt,
}

var AvailableProductsForRentingResponse = [] response.AvailableProductForRentingResponse{
	AvailableProductForRentingResponse1,
	AvailableProductForRentingResponse2,
}