package service

import (
	. "../model"
	. "../repository"
)

type ICityService interface {
	FindAll() []City
}

type CityService struct {
	ICityRepository
}

func (cityService *CityService) FindAll() []City {
	return cityService.FindAllCities()
}