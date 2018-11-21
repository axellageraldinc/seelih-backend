package mapper

import (
	"../model"
	"../model/response"
)

func ToCategoryResponseList(categories []model.Category) []response.CategoryResponse {
	var categoryResponseList []response.CategoryResponse
	for index := range categories {
		availableProductForRentingResponse := response.CategoryResponse{
			Id:                 categories[index].ID,
			Name:               categories[index].Name,
		}
		categoryResponseList = append(categoryResponseList, availableProductForRentingResponse)
	}
	return categoryResponseList
}