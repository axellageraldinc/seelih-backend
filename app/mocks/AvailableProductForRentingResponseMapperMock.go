package mocks

import (
	. "../model"
	. "../model/response"
	"github.com/stretchr/testify/mock"
)

type IAvailableProductForRentingResponseMapper struct {
	mock.Mock
}

func (mock *IAvailableProductForRentingResponseMapper) ToAvailableProductForRentingResponseList(products []Product) []AvailableProductForRentingResponse {
	args := mock.Called(products)
	return args.Get(0).([]AvailableProductForRentingResponse)
}