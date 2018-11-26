package mocks

import (
	. "../model"
	. "../model/response"
	"github.com/stretchr/testify/mock"
)

type IUploadedProductResponseMapper struct {
	mock.Mock
}

func (mock *IUploadedProductResponseMapper) ToUploadedProductResponseList(products []Product) []UploadedProductResponse {
	args := mock.Called(products)
	return args.Get(0).([]UploadedProductResponse)
}