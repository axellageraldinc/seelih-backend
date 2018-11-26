package mocks

import (
	. "../model"
	"github.com/stretchr/testify/mock"
)

type IUserRepository struct {
	mock.Mock
}

func (mock *IUserRepository) SaveUser(user User) bool {
	args := mock.Called(user)
	return args.Bool(0)
}

func (mock *IUserRepository) FindUserByEmail(email string) User {
	args := mock.Called(email)
	return args.Get(0).(User)
}

func (mock *IUserRepository) FindUserById(id int) User {
	args := mock.Called(id)
	return args.Get(0).(User)
}

func (mock *IUserRepository) DoesUserEmailExist(email string) bool {
	args := mock.Called(email)
	return args.Bool(0)
}

func (mock *IUserRepository) DoesUserIdExist(id int) bool {
	args := mock.Called(id)
	return args.Bool(0)
}