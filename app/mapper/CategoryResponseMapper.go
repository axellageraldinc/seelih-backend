package mapper

import (
	. "../model"
	. "../model/response"
)

type ICategoryResponseMapper interface {
	ToCategoryResponseList([]Category) []CategoryResponse
}

type CategoryResponseMapper struct {}

func (categoryResponseMapper *CategoryResponseMapper) ToCategoryResponseList(categories []Category) []CategoryResponse {
	var categoryResponseList []CategoryResponse
	for index := range categories {
		availableProductForRentingResponse := CategoryResponse{
			Id:                 categories[index].ID,
			Name:               categories[index].Name,
		}
		categoryResponseList = append(categoryResponseList, availableProductForRentingResponse)
	}
	return categoryResponseList
}