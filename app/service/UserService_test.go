package service

import (
	"../mocks"
	"../model"
	"../util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserService_GetUserData_Success(t *testing.T) {
	userRepository := new(mocks.IUserRepository)

	userRepository.On("DoesUserIdExist", 1).Return(true)
	userRepository.On("FindUserById", 1).Return(util.User1)

	userService := UserService{
		IUserRepository: userRepository,
	}

	expectedResult := util.User1
	expectedErrorCode := 0

	actualResult, actualErrorCode := userService.GetUserData(1)

	assert.NotNil(t, actualResult)
	assert.NotNil(t, actualErrorCode)
	assert.Equal(t, expectedErrorCode, int(actualErrorCode))
	assert.Equal(t, expectedResult, actualResult)
}

func TestUserService_GetUserData_Failed_UserIdNotFound(t *testing.T) {
	userRepository := new(mocks.IUserRepository)

	userRepository.On("DoesUserIdExist", 2).Return(false)

	userService := UserService{
		IUserRepository: userRepository,
	}

	expectedResult := model.User{}
	expectedErrorCode := model.LOGIN_FAILED

	actualResult, actualErrorCode := userService.GetUserData(2)

	assert.NotNil(t, actualResult)
	assert.NotNil(t, actualErrorCode)
	assert.Equal(t, expectedErrorCode, int(actualErrorCode))
	assert.Equal(t, expectedResult, actualResult)
}

func TestUserService_Login_Success(t *testing.T) {
	userRepository := new(mocks.IUserRepository)

	userRepository.On("DoesUserEmailExist", util.LoginRequest1.Email).Return(true)
	userRepository.On("FindUserByEmail", util.LoginRequest1.Email).Return(util.User1)
	hashedPasswordFromDatabase := util.User1.Password
	isPasswordTrue := comparePasswords(hashedPasswordFromDatabase, ConvertPlainPasswordToByte(util.LoginRequest1.Password))

	userService := UserService{
		IUserRepository: userRepository,
	}

	expectedResult := util.User1
	expectedErrorCode := 0

	actualResult, actualErrorCode := userService.Login(util.LoginRequest1)

	assert.NotNil(t, actualResult)
	assert.NotNil(t, actualErrorCode)
	assert.True(t, isPasswordTrue)
	assert.Equal(t, expectedErrorCode, int(actualErrorCode))
	assert.Equal(t, expectedResult, actualResult)
}

func TestUserService_Login_Failed_PasswordFailed(t *testing.T) {
	userRepository := new(mocks.IUserRepository)

	userRepository.On("DoesUserEmailExist", util.LoginRequest2.Email).Return(true)
	userRepository.On("FindUserByEmail", util.LoginRequest2.Email).Return(model.User{})
	hashedPasswordFromDatabase := util.User2.Password
	isPasswordTrue := comparePasswords(hashedPasswordFromDatabase, ConvertPlainPasswordToByte(util.LoginRequest2.Password))

	userService := UserService{
		IUserRepository: userRepository,
	}

	expectedResult := model.User{}
	expectedErrorCode := model.LOGIN_FAILED

	actualResult, actualErrorCode := userService.Login(util.LoginRequest2)

	assert.NotNil(t, actualResult)
	assert.NotNil(t, actualErrorCode)
	assert.False(t, isPasswordTrue)
	assert.Equal(t, expectedErrorCode, int(actualErrorCode))
	assert.Equal(t, expectedResult, actualResult)
}

func TestUserService_Login_Failed_EmailNotExists(t *testing.T) {
	userRepository := new(mocks.IUserRepository)

	userRepository.On("DoesUserEmailExist", util.LoginRequest2.Email).Return(false)

	userService := UserService{
		IUserRepository: userRepository,
	}

	expectedResult := model.User{}
	expectedErrorCode := model.LOGIN_FAILED

	actualResult, actualErrorCode := userService.Login(util.LoginRequest2)

	assert.NotNil(t, actualResult)
	assert.NotNil(t, actualErrorCode)
	assert.Equal(t, expectedErrorCode, int(actualErrorCode))
	assert.Equal(t, expectedResult, actualResult)
}

func TestUserService_Register_Failed_EmailAlreadyExists(t *testing.T) {
	userRepository := new(mocks.IUserRepository)

	userRepository.On("DoesUserEmailExist", util.RegisterRequest1.Email).Return(true)

	userService := UserService{
		IUserRepository: userRepository,
	}

	expectedResult := false
	expectedErrorCode := model.REGISTER_FAILED_EMAIL_ALREADY_EXISTS

	actualResult, actualErrorCode := userService.Register(util.RegisterRequest1)

	assert.NotNil(t, actualResult)
	assert.NotNil(t, actualErrorCode)
	assert.Equal(t, expectedErrorCode, int(actualErrorCode))
	assert.Equal(t, expectedResult, actualResult)
}

func TestUserService_Register_Success(t *testing.T) {
	userRepository := new(mocks.IUserRepository)

	var registerRequest = util.RegisterRequest1

	userRepository.On("DoesUserEmailExist", registerRequest.Email).Return(false)
	userRepository.On("SaveUser", mock.AnythingOfType("User")).Return(true)

	userService := UserService{
		IUserRepository: userRepository,
	}

	expectedResult := true
	expectedErrorCode := 0

	actualResult, actualErrorCode := userService.Register(registerRequest)

	assert.NotNil(t, actualResult)
	assert.NotNil(t, actualErrorCode)
	assert.Equal(t, expectedErrorCode, int(actualErrorCode))
	assert.Equal(t, expectedResult, actualResult)
}

func TestUserService_Register_Failed_WontSaveToDatabase(t *testing.T) {
	userRepository := new(mocks.IUserRepository)

	userRepository.On("DoesUserEmailExist", util.RegisterRequest1.Email).Return(false)
	userRepository.On("SaveUser", mock.AnythingOfType("User")).Return(false)

	userService := UserService{
		IUserRepository: userRepository,
	}

	expectedResult := false
	expectedErrorCode := model.REGISTER_FAILED_WONT_SAVE_TO_DATABASE

	actualResult, actualErrorCode := userService.Register(util.RegisterRequest1)

	assert.NotNil(t, actualResult)
	assert.NotNil(t, actualErrorCode)
	assert.Equal(t, expectedErrorCode, int(actualErrorCode))
	assert.Equal(t, expectedResult, actualResult)
}