package mocks

import (
	. "../model"
	"github.com/stretchr/testify/mock"
)

type ICityService struct {
	mock.Mock
}

func (mock *ICityService) FindAll() []City {
	args := mock.Called()
	return args.Get(0).([]City)
}