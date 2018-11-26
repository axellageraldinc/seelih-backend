package mocks

import (
	. "../model"
	"github.com/stretchr/testify/mock"
)

type ICategoryService struct {
	mock.Mock
}

func (mock *ICategoryService) FindAll() []Category {
	args := mock.Called()
	return args.Get(0).([]Category)
}