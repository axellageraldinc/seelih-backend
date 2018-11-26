package mocks

import (
	. "../model"
	. "../model/request"
	"github.com/stretchr/testify/mock"
)

type IUserService struct {
	mock.Mock
}

func (mock *IUserService) Register(request RegisterRequest) (bool, uint) {
	args := mock.Called(request)
	return args.Bool(0), uint(args.Int(1))
}

func (mock *IUserService) Login(request LoginRequest) (User, uint) {
	args := mock.Called(request)
	return args.Get(0).(User), uint(args.Int(1))
}

func (mock *IUserService) GetUserData(userId int) (User, uint) {
	args := mock.Called(userId)
	return args.Get(0).(User), uint(args.Int(1))
}

func (mock *IUserService) HashAndSalt(bytePlainPassword []byte) string {
	args := mock.Called(bytePlainPassword)
	return args.String(0)
}

func (mock *IUserService) ConvertPlainPasswordToByte(plainPassword string) []byte {
	args := mock.Called(plainPassword)
	return args.Get(0).([]byte)
}

func (mock *IUserService) comparePasswords(hashedPassword string, plainPassword []byte) bool {
	args := mock.Called(hashedPassword, plainPassword)
	return args.Bool(0)
}