package mapper

import (
	"../model"
	"../model/response"
)

func ToCityResponseList(cities []model.City) []response.CityResponse {
	var cityResponseList []response.CityResponse
	for index := range cities {
		cityResponse := response.CityResponse{
			Id:   cities[index].ID,
			Name: cities[index].Name,
		}
		cityResponseList = append(cityResponseList, cityResponse)
	}
	return cityResponseList
}
