package util

import (
	"../model/response"
)

var CategoryResponse1 = response.CategoryResponse{
	Name: "category1",
	Id: 1,
}
var CategoryResponse2 = response.CategoryResponse{
	Name: "category2",
	Id: 2,
}
var CategoriesResponse = []response.CategoryResponse{
	CategoryResponse1,
	CategoryResponse2,
}