package mocks

import (
	. "../model"
	. "../model/response"
	"github.com/stretchr/testify/mock"
)

type ICityResponseMapper struct {
	mock.Mock
}

func (mock *ICityResponseMapper) ToCityResponseList(cities []City) []CityResponse {
	args := mock.Called(cities)
	return args.Get(0).([]CityResponse)
}