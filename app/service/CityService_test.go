package service

import (
	"../mocks"
	"../util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCityService_FindAll_Success(t *testing.T) {
	cityRepository := new(mocks.ICityRepository)

	cityRepository.On("FindAllCities").Return(util.Cities)

	cityService := CityService{
		ICityRepository: cityRepository,
	}

	expectedResult := util.Cities

	actualResult := cityService.FindAll()

	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult[0].Name, actualResult[0].Name)
	assert.Equal(t, expectedResult[1].Name, actualResult[1].Name)
}
