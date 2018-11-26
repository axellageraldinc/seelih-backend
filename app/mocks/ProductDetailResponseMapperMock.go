package mocks

import (
	. "../model"
	. "../model/response"
	"github.com/stretchr/testify/mock"
)

type IProductDetailResponseMapper struct {
	mock.Mock
}

func (mock *IProductDetailResponseMapper) ToProductDetailResponse(product Product) ProductDetailResponse {
	args := mock.Called(product)
	return args.Get(0).(ProductDetailResponse)
}