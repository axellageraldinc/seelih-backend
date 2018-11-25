package repository

import (
	. "../helper"
	. "../model"
)

type ICityRepository interface {
	FindAllCities() []City
}

type CityRepository struct {
	IDatabaseConnectionHelper
}

func (cityRepository *CityRepository) FindAllCities() []City {
	var cities []City
	db := cityRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Find(&cities)
	return cities
}