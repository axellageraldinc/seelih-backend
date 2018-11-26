package util

import (
	"../model/response"
)

var CityResponse1 = response.CityResponse{
	Id: 1,
	Name: "CityResponse 1",
}
var CityResponse2 = response.CityResponse{
	Id: 2,
	Name: "CityResponse 2",
}
var CitiesResponse = []response.CityResponse{
	CityResponse1,
	CityResponse2,
}