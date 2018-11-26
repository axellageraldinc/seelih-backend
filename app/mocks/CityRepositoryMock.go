package mocks

import (
	"github.com/stretchr/testify/mock"
	. "../model"
)

type ICityRepository struct {
	mock.Mock
}

func (mock *ICityRepository) FindAllCities() []City {
	args := mock.Called()
	return args.Get(0).([]City)
}