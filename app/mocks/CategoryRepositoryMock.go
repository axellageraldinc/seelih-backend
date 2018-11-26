package mocks

import (
	. "../model"
	"github.com/stretchr/testify/mock"
)

type ICategoryRepository struct {
	mock.Mock
}

func (mock *ICategoryRepository) FindAllCategories() []Category {
	args := mock.Called()
	return args.Get(0).([]Category)
}

func (mock *ICategoryRepository) DoesCategoryIdExist(id int) bool {
	args := mock.Called(id)
	return args.Bool(0)
}