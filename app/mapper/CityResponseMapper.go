package mapper

import (
	. "../model"
	. "../model/response"
)

type ICityResponseMapper interface {
	ToCityResponseList([]City) []CityResponse
}

type CityResponseMapper struct {}

func (cityResponseMapper *CityResponseMapper) ToCityResponseList(cities []City) []CityResponse {
	var cityResponseList []CityResponse
	for index := range cities {
		cityResponse := CityResponse{
			Id:   cities[index].ID,
			Name: cities[index].Name,
		}
		cityResponseList = append(cityResponseList, cityResponse)
	}
	return cityResponseList
}
