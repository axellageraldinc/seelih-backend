package controller

import (
	"../mocks"
	"../model"
	"../model/response"
	"../util"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUserController_GetUserDataHandler_Success(t *testing.T) {
	userService := new(mocks.IUserService)
	userLoginMapper := new(mocks.IUserLoginMapper)

	userService.On("GetUserData", 1).Return(util.User1, 0)
	userLoginMapper.On("ToUserLoginDetail", util.User1).Return(util.UserResponse1)

	userController := UserController{
		IUserLoginMapper: userLoginMapper,
		IUserService: userService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/users/1", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/{userId}", userController.GetUserDataHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: util.UserResponse1,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.NotNil(t, actualResponse.Data)
	assert.NotEmpty(t, actualResponse.Data)
	var expectedData = expectedResponse.Data.(response.UserResponse)
	var actualData = actualResponse.Data.(interface{})
	assert.Equal(t, float64(expectedData.Id) , actualData.(map[string]interface{})["Id"])
}

func TestUserController_GetUserDataHandler_Failed_ErrorCodeExists(t *testing.T) {
	userService := new(mocks.IUserService)
	userLoginMapper := new(mocks.IUserLoginMapper)

	userService.On("GetUserData", 999).Return(model.User{}, model.LOGIN_FAILED)

	userController := UserController{
		IUserLoginMapper: userLoginMapper,
		IUserService: userService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/users/999", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/{userId}", userController.GetUserDataHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.LOGIN_FAILED,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestUserController_LoginHandler_Success(t *testing.T) {
	userService := new(mocks.IUserService)
	userLoginMapper := new(mocks.IUserLoginMapper)

	userService.On("Login", util.LoginRequest1).Return(util.User1, 0)
	userLoginMapper.On("ToUserLoginDetail", util.User1).Return(util.UserResponse1)

	userController := UserController{
		IUserLoginMapper: userLoginMapper,
		IUserService: userService,
	}

	jsonLoginRequest, _ := json.Marshal(util.LoginRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/users/login", bytes.NewBuffer(jsonLoginRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/login", userController.LoginHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: util.UserResponse1,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.NotNil(t, actualResponse.Data)
	assert.NotEmpty(t, actualResponse.Data)
	var expectedData = expectedResponse.Data.(response.UserResponse)
	var actualData = actualResponse.Data.(interface{})
	assert.Equal(t, float64(expectedData.Id) , actualData.(map[string]interface{})["Id"])
}

func TestUserController_LoginHandler_Failed_ErrorCodeExists(t *testing.T) {
	userService := new(mocks.IUserService)
	userLoginMapper := new(mocks.IUserLoginMapper)

	userService.On("Login", util.LoginRequest1).Return(model.User{}, model.LOGIN_FAILED)

	userController := UserController{
		IUserLoginMapper: userLoginMapper,
		IUserService: userService,
	}

	jsonLoginRequest, _ := json.Marshal(util.LoginRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/users/login", bytes.NewBuffer(jsonLoginRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/login", userController.LoginHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.LOGIN_FAILED,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestUserController_RegisterHandler_Success(t *testing.T) {
	userService := new(mocks.IUserService)
	userLoginMapper := new(mocks.IUserLoginMapper)

	userService.On("Register", util.RegisterRequest1).Return(true, 0)

	userController := UserController{
		IUserLoginMapper: userLoginMapper,
		IUserService: userService,
	}

	jsonRegisterRequest, _ := json.Marshal(util.RegisterRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/users/register", bytes.NewBuffer(jsonRegisterRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/register", userController.RegisterHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestUserController_RegisterHandler_Failed_ErrorCodeExists(t *testing.T) {
	userService := new(mocks.IUserService)
	userLoginMapper := new(mocks.IUserLoginMapper)

	userService.On("Register", util.RegisterRequest1).Return(false, model.REGISTER_FAILED_EMAIL_ALREADY_EXISTS)

	userController := UserController{
		IUserLoginMapper: userLoginMapper,
		IUserService: userService,
	}

	jsonRegisterRequest, _ := json.Marshal(util.RegisterRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/users/register", bytes.NewBuffer(jsonRegisterRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/register", userController.RegisterHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.REGISTER_FAILED_EMAIL_ALREADY_EXISTS,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}