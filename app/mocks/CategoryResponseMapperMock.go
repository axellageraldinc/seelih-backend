package mocks

import (
	. "../model"
	. "../model/response"
	"github.com/stretchr/testify/mock"
)

type ICategoryResponseMapper struct {
	mock.Mock
}

func (mock *ICategoryResponseMapper) ToCategoryResponseList(categories []Category) []CategoryResponse {
	args := mock.Called(categories)
	return args.Get(0).([]CategoryResponse)
}