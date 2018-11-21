package mapper

import (
	. "../model"
	. "../model/response"
)

func ToAvailableProductForRentingResponseList(products []Product) []AvailableProductForRentingResponse {
	var availableProductForRentingResponseList []AvailableProductForRentingResponse
	for index := range products {
		availableProductForRentingResponse := AvailableProductForRentingResponse{
			Id:                 products[index].ID,
			Name:               products[index].Name,
			PricePerItemPerDay: products[index].PricePerItemPerDay,
			UploadedTime:       products[index].CreatedAt,
			ImageUrl: IMAGE_URL_PREFIX + products[index].ImageName,
		}
		availableProductForRentingResponseList = append(availableProductForRentingResponseList, availableProductForRentingResponse)
	}
	return availableProductForRentingResponseList
}