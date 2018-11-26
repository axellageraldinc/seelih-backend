package mocks

import (
	. "../model"
	. "../model/response"
	"github.com/stretchr/testify/mock"
)

type IUserLoginMapper struct {
	mock.Mock
}

func (mock *IUserLoginMapper) ToUserLoginDetail(user User) UserResponse {
	args := mock.Called(user)
	return args.Get(0).(UserResponse)
}