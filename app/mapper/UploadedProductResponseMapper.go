package mapper

import (
	. "../model"
	. "../model/response"
)

func ToUploadedProductResponseList(products []Product) []UploadedProductResponse {
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