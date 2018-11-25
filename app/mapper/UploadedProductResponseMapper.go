package mapper

import (
	. "../model"
	. "../model/response"
)

type IUploadedProductResponseMapper interface {
	ToUploadedProductResponseList(products []Product) []UploadedProductResponse
}

type UploadedProductResponseMapper struct {}

func (uploadedProductResponseMapper *UploadedProductResponseMapper) ToUploadedProductResponseList(products []Product) []UploadedProductResponse {
	var uploadedProductResponseList []UploadedProductResponse
	for index := range products {
		uploadedProductResponse := UploadedProductResponse{
			Id:                 products[index].ID,
			Name:               products[index].Name,
			PricePerItemPerDay: products[index].PricePerItemPerDay,
			UploadedTime:       products[index].CreatedAt,
			ImageUrl: IMAGE_URL_PREFIX + products[index].ImageName,
		}
		uploadedProductResponseList = append(uploadedProductResponseList, uploadedProductResponse)
	}
	return uploadedProductResponseList
}