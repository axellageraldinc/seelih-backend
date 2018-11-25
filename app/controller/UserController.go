package controller

import (
	. "../helper"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	. "../mapper"
	. "../model/request"
	. "../model/response"
	. "../service"
	"github.com/kataras/golog"
)

type UserController struct {
	IUserService
	IUserLoginMapper
}

func (userController *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/register")

	var registerRequest RegisterRequest
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&registerRequest)

	isSuccess, errorCode := userController.Register(registerRequest)
	if isSuccess {
		response = OK(nil)
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}

func (userController *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/login")

	var loginRequest LoginRequest
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&loginRequest)

	user, errorCode := userController.Login(loginRequest)
	if errorCode == 0 {
		response = OK(userController.ToUserLoginDetail(user))
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}

func (userController *UserController) GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/{userId}")

	parameters := mux.Vars(r)
	userId, _ := strconv.ParseInt(parameters["userId"], 10, 32)

	var response WebResponse

	user, errorCode := userController.GetUserData(int(userId))

	if errorCode == 0 {
		response = OK(userController.ToUserLoginDetail(user))
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}
